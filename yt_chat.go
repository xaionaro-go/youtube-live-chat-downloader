package YtChat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type SubMenuItems struct {
	Title        string
	Continuation struct {
		ReloadContinuationData struct {
			Continuation string
		}
	}
}

type InnerTubeContext struct {
	Client struct {
		Hl               string `json:"hl"`
		Gl               string `json:"gl"`
		RemoteHost       string `json:"remoteHost"`
		DeviceMake       string `json:"deviceMake"`
		DeviceModel      string `json:"deviceModel"`
		VisitorData      string `json:"visitorData"`
		UserAgent        string `json:"userAgent"`
		ClientName       string `json:"clientName"`
		ClientVersion    string `json:"clientVersion"`
		OsName           string `json:"osName"`
		OsVersion        string `json:"osVersion"`
		OriginalUrl      string `json:"originalUrl"`
		Platform         string `json:"platform"`
		ClientFormFactor string `json:"clientFormFactor"`
		ConfigInfo       struct {
			AppInstallData string `json:"appInstallData"`
		} `json:"configInfo"`
	} `json:"client"`
}

type YtCfg struct {
	INNERTUBE_API_KEY             string
	INNERTUBE_CONTEXT             InnerTubeContext
	INNERTUBE_CONTEXT_CLIENT_NAME string
	INNERTUBE_CLIENT_VERSION      string
	ID_TOKEN                      string
}

type Context struct {
	Context      InnerTubeContext `json:"context"`
	Continuation string           `json:"continuation"`
}

type ContinuationChat struct {
	TimedContinuationData struct {
		Continuation string `json:"continuation"`
		TimeoutMs    int    `json:"timeoutMs"`
	} `json:"timedContinuationData"`
	InvalidationContinuationData struct {
		Continuation string `json:"continuation"`
		TimeoutMs    int    `json:"timeoutMs"`
	} `json:"invalidationContinuationData"`
}

type Thumbnail struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type Actions struct {
	AddChatItemAction struct {
		Item struct {
			LiveChatTextMessageRenderer struct {
				Message struct {
					Runs []Runs `json:"runs"`
				} `json:"message"`
				AuthorName struct {
					SimpleText string `json:"simpleText"`
				}
				AuthorPhoto struct {
					Thumbnails []Thumbnail `json:"thumbnails"`
				} `json:"authorPhoto"`
				TimestampUsec           string `json:"timestampUsec"`
				AuthorExternalChannelId string `json:"authorExternalChannelId"`
			} `json:"liveChatTextMessageRenderer"`
		} `json:"item"`
	} `json:"addChatItemAction"`
}

type Runs struct {
	Text  string `json:"text,omitempty"`
	Emoji struct {
		EmojiId       string `json:"emojiId"`
		IsCustomEmoji bool   `json:"isCustomEmoji,omitempty"`
		Image         struct {
			Thumbnails []Thumbnail
		}
	} `json:"emoji,omitempty"`
}

type ChatMessagesResponse struct {
	ContinuationContents struct {
		LiveChatContinuation struct {
			Actions       []Actions          `json:"actions"`
			Continuations []ContinuationChat `json:"continuations"`
		} `json:"liveChatContinuation"`
	} `json:"continuationContents"`
}

type ChatMessage struct {
	AuthorID   string
	AuthorName string
	Message    string
	Timestamp  time.Time
	Thumbnails []Thumbnail
}

type InitialData struct {
	Contents struct {
		TwoColumnWatchNextResults struct {
			ConversationBar struct {
				LiveChatRenderer struct {
					Header struct {
						LiveChatHeaderRenderer struct {
							ViewSelector struct {
								SortFilterSubMenuRenderer struct {
									SubMenuItems []SubMenuItems `json:"subMenuItems"`
								}
							}
						}
					}
				}
			}
		}
	}
}

var (
	LIVE_CHAT_URL = `https://www.youtube.com/youtubei/v1/live_chat/get_%s?key=%s`
	// Google would sometimes ask you to solve a CAPTCHA before accessing it's websites
	// or ask for your CONSENT if you are an EU user
	// You can add those cookies here.
	customCookies     []*http.Cookie
	ErrLiveStreamOver error = errors.New("live stream over")
	ErrStreamNotLive  error = errors.New("stream not live")
)

const (
	API_TYPE           = "live_chat"
	YT_CFG_REGEX       = `ytcfg\.set\s*\(\s*({.+?})\s*\)\s*;`
	INITIAL_DATA_REGEX = `(?:window\s*\[\s*["\']ytInitialData["\']\s*\]|ytInitialData)\s*=\s*({.+?})\s*;\s*(?:var\s+meta|</script|\n)`
)

func regexSearch(regex string, str string) []string {
	r, _ := regexp.Compile(regex)
	matches := r.FindAllString(str, -1)
	return matches
}

func parseMicroSeconds(timeStampStr string) time.Time {
	tm, _ := strconv.ParseInt(timeStampStr, 10, 64)
	tm = tm / 1000
	sec := tm / 1000
	msec := tm % 1000
	return time.Unix(sec, msec*int64(time.Millisecond))
}

func fetchChatMessages(initialContinuationInfo string, ytCfg YtCfg) ([]ChatMessage, string, int, error) {
	apiKey := ytCfg.INNERTUBE_API_KEY
	continuationUrl := fmt.Sprintf(LIVE_CHAT_URL, API_TYPE, apiKey)
	innertubeContext := ytCfg.INNERTUBE_CONTEXT

	context := Context{innertubeContext, initialContinuationInfo}
	b, _ := json.Marshal(context)
	var jsonData = []byte(b)
	request, error := http.NewRequest("POST", continuationUrl, bytes.NewBuffer(jsonData))
	if error != nil {
		return nil, "", 0, error
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return nil, "", 0, error
	}
	if response.StatusCode != 200 {
		return nil, "", 0, fmt.Errorf("some error fetching chat messages status code: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", 0, err
	}
	response.Body.Close()
	var chatMsgResp ChatMessagesResponse
	json.Unmarshal([]byte(string(body)), &chatMsgResp)
	actions := chatMsgResp.ContinuationContents.LiveChatContinuation.Actions

	chatMessages := []ChatMessage{}
	for _, action := range actions {
		liveChatTextMessageRenderer := action.AddChatItemAction.Item.LiveChatTextMessageRenderer
		// Each chat message is separated into multiple runs.
		// Iterate through all runs and generate the chat message.
		runs := liveChatTextMessageRenderer.Message.Runs
		if len(runs) > 0 {
			chatMessage := ChatMessage{
				AuthorID:   liveChatTextMessageRenderer.AuthorExternalChannelId,
				AuthorName: liveChatTextMessageRenderer.AuthorName.SimpleText,
				Timestamp:  parseMicroSeconds(liveChatTextMessageRenderer.TimestampUsec),
				Thumbnails: liveChatTextMessageRenderer.AuthorPhoto.Thumbnails,
			}
			text := ""
			for _, run := range runs {
				if run.Text != "" {
					text += run.Text
				} else {
					if run.Emoji.IsCustomEmoji {
						numberOfThumbnails := len(run.Emoji.Image.Thumbnails)
						// Youtube chat has custom emojis which
						// are small PNG images and cannot be displayed as text.
						//
						// These custom emojis are available with their image url.
						//
						// Adding some whitespace around custom image URLs
						// without the whitespace it would be difficult to parse these URLs
						if numberOfThumbnails > 0 && numberOfThumbnails == 2 {
							text += " " + run.Emoji.Image.Thumbnails[1].URL + " "
						} else if numberOfThumbnails == 1 {
							text += " " + run.Emoji.Image.Thumbnails[0].URL + " "
						}
					} else {
						text += run.Emoji.EmojiId
					}
				}
			}
			chatMessage.Message = text
			chatMessages = append(chatMessages, chatMessage)
		}
	}
	// No continuation returned from youtube, Stream has ended.
	if len(chatMsgResp.ContinuationContents.LiveChatContinuation.Continuations) == 0 {
		return nil, "", 0, ErrLiveStreamOver
	}
	// extract continuation and timeout received from response
	timeoutMs := 5
	continuations := chatMsgResp.ContinuationContents.LiveChatContinuation.Continuations[0]
	if continuations.TimedContinuationData.Continuation == "" {
		initialContinuationInfo = continuations.InvalidationContinuationData.Continuation
		timeoutMs = continuations.InvalidationContinuationData.TimeoutMs
	} else {
		initialContinuationInfo = continuations.TimedContinuationData.Continuation
		timeoutMs = continuations.TimedContinuationData.TimeoutMs
	}
	return chatMessages, initialContinuationInfo, timeoutMs, nil
}

func ParseInitialData(videoUrl string) (string, YtCfg, error) {
	req, err := http.NewRequest("GET", videoUrl, nil)
	if err != nil {
		return "", YtCfg{}, err
	}

	for _, cookie := range customCookies {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", YtCfg{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Print(videoUrl +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return "", YtCfg{}, err
	}

	intArr, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", YtCfg{}, err
	}

	html := string(intArr)

	// TODO ::  work on regex
	initialDataArr := regexSearch(INITIAL_DATA_REGEX, html)
	initialData := strings.Trim(string(initialDataArr[0]), "ytInitialData = ") // TODO: fix this misuse of Trim
	initialData = strings.Trim(initialData, ";</script")
	ytCfg := regexSearch(YT_CFG_REGEX, html)[0]
	ytCfg = strings.Trim(ytCfg, "ytcfg.set(")
	ytCfg = strings.Trim(ytCfg, ");")

	var _ytCfg YtCfg
	json.Unmarshal([]byte(ytCfg), &_ytCfg)

	var _initialData InitialData
	json.Unmarshal([]byte(initialData), &_initialData)

	subMenuItems := _initialData.Contents.TwoColumnWatchNextResults.ConversationBar.LiveChatRenderer.Header.LiveChatHeaderRenderer.ViewSelector.SortFilterSubMenuRenderer.SubMenuItems
	if len(subMenuItems) == 0 {
		return "", YtCfg{}, ErrStreamNotLive
	}
	initialContinuationInfo := subMenuItems[1].Continuation.ReloadContinuationData.Continuation
	return initialContinuationInfo, _ytCfg, nil
}

func FetchContinuationChat(continuation string, ytCfg YtCfg) ([]ChatMessage, string, error) {
	chatMessages, continuation, timeoutMs, error := fetchChatMessages(continuation, ytCfg)
	if error != nil {
		return nil, "", error
	}
	// Sleep for timeoutMs milliseconds sent by the server
	if timeoutMs > 0 {
		time.Sleep(time.Duration(timeoutMs) * time.Millisecond)
	} else {
		time.Sleep(time.Second * 5)
	}
	return chatMessages, continuation, nil
}

func AddCookies(cookies []*http.Cookie) {
	customCookies = cookies
}
