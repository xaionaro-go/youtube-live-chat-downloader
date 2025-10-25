package YtChat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

type MainAppWebInfo struct {
	GraftURL                  string `json:"graftUrl"`
	WebDisplayMode            string `json:"webDisplayMode"`
	IsWebNativeShareAvailable bool   `json:"isWebNativeShareAvailable"`
}

type ConfigInfo struct {
	AppInstallData string `json:"appInstallData"`
}

type InnerTubeContext struct {
	Client struct {
		Hl                 string         `json:"hl"`
		Gl                 string         `json:"gl"`
		RemoteHost         string         `json:"remoteHost"`
		DeviceMake         string         `json:"deviceMake"`
		DeviceModel        string         `json:"deviceModel"`
		VisitorData        string         `json:"visitorData"`
		UserAgent          string         `json:"userAgent"`
		ClientName         string         `json:"clientName"`
		ClientVersion      string         `json:"clientVersion"`
		OsName             string         `json:"osName"`
		OsVersion          string         `json:"osVersion"`
		OriginalURL        string         `json:"originalUrl"`
		ScreenPixelDensity int            `json:"screenPixelDensity"`
		Platform           string         `json:"platform"`
		ClientFormFactor   string         `json:"clientFormFactor"`
		ConfigInfo         ConfigInfo     `json:"configInfo"`
		ScreenDensityFloat int            `json:"screenDensityFloat"`
		UserInterfaceTheme string         `json:"userInterfaceTheme"`
		TimeZone           string         `json:"timeZone"`
		BrowserName        string         `json:"browserName"`
		BrowserVersion     string         `json:"browserVersion"`
		AcceptHeader       string         `json:"acceptHeader"`
		DeviceExperimentID string         `json:"deviceExperimentId"`
		RolloutToken       string         `json:"rolloutToken"`
		ScreenWidthPoints  int            `json:"screenWidthPoints"`
		ScreenHeightPoints int            `json:"screenHeightPoints"`
		UtcOffsetMinutes   int            `json:"utcOffsetMinutes"`
		MainAppWebInfo     MainAppWebInfo `json:"mainAppWebInfo"`
	} `json:"client"`
	User struct {
		LockedSafetyMode bool `json:"lockedSafetyMode"`
	} `json:"user"`
	Request struct {
		UseSsl                  bool  `json:"useSsl"`
		InternalExperimentFlags []any `json:"internalExperimentFlags"`
		ConsistencyTokenJars    []any `json:"consistencyTokenJars"`
	} `json:"request"`
	ClickTracking struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
	} `json:"clickTracking"`
	AdSignalsInfo struct {
		Params []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"params"`
	} `json:"adSignalsInfo"`
}

type YtCfg struct {
	INNERTUBE_API_KEY             string
	INNERTUBE_CONTEXT             InnerTubeContext
	INNERTUBE_CONTEXT_CLIENT_NAME any
	INNERTUBE_CLIENT_VERSION      string
	ID_TOKEN                      string
}

type WebClientInfo struct {
	IsDocumentHidden bool `json:"isDocumentHidden"`
}

type Context struct {
	Context       InnerTubeContext `json:"context"`
	Continuation  string           `json:"continuation"`
	WebClientInfo WebClientInfo    `json:"webClientInfo"`
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

var (
	LIVE_CHAT_URL = `https://www.youtube.com/youtubei/v1/live_chat/get_live_chat?prettyPrint=false`
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
	continuationUrl := LIVE_CHAT_URL
	innertubeContext := ytCfg.INNERTUBE_CONTEXT

	context := Context{
		Context:      innertubeContext,
		Continuation: initialContinuationInfo,
		WebClientInfo: WebClientInfo{
			IsDocumentHidden: false,
		},
	}
	//context.Context.Client.OriginalURL = fmt.Sprintf("https://www.youtube.com/live_chat?continuation=%s", initialContinuationInfo)
	//context.Context.Client.MainAppWebInfo.GraftURL = fmt.Sprintf("https://www.youtube.com/live_chat?continuation=%s", initialContinuationInfo)
	context.Context.Client.MainAppWebInfo.WebDisplayMode = "WEB_DISPLAY_MODE_BROWSER"

	jsonData, err := json.Marshal(context)
	if err != nil {
		return nil, "", 0, fmt.Errorf("error marshalling context %v: %w", context, err)
	}
	request, err := http.NewRequest("POST", continuationUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, "", 0, fmt.Errorf("error creating POST request to %s with data %s: %w", continuationUrl, string(jsonData), err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, "", 0, fmt.Errorf("error performing POST request to %s with data %s: %w", continuationUrl, string(jsonData), err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, "", 0, fmt.Errorf("error reading response body from %s with data %s: %w", continuationUrl, string(jsonData), err)
	}

	if response.StatusCode != 200 {
		return nil, "", 0, fmt.Errorf("error fetching chat messages from %s with data %s: status code: %d: %s", continuationUrl, string(jsonData), response.StatusCode, body)
	}

	var chatMsgResp ChatMessagesResponse
	json.Unmarshal(body, &chatMsgResp)
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

func GetYTDataFromURL(urlString string) ([]byte, []byte, error) {
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create a request to '%s': %w", urlString, err)
	}

	for _, cookie := range customCookies {
		req.AddCookie(cookie)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to perform a request to '%s': %w", urlString, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := fmt.Errorf("received a non-200 code from '%s': %d", urlString, resp.StatusCode)
		return nil, nil, err
	}

	intArr, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read the body of the response from '%s': %w", urlString, err)
	}

	html := string(intArr)

	// TODO ::  work on regex
	initialDataArr := regexSearch(INITIAL_DATA_REGEX, html)
	initialData := strings.Trim(string(initialDataArr[0]), "ytInitialData = ") // TODO: fix this misuse of Trim
	initialData = strings.Trim(initialData, ";</script")
	ytCfg := regexSearch(YT_CFG_REGEX, html)[0]
	ytCfg = strings.Trim(ytCfg, "ytcfg.set(")
	ytCfg = strings.Trim(ytCfg, ");")

	return []byte(initialData), []byte(ytCfg), nil
}

func ParseInitialData(videoUrl string) (string, YtCfg, error) {
	initialData, ytCfg, err := GetYTDataFromURL(videoUrl)
	if err != nil {
		return "", YtCfg{}, fmt.Errorf("unable to get ytcfg: %w", err)
	}

	var _ytCfg YtCfg
	err = json.Unmarshal(ytCfg, &_ytCfg)
	if err != nil {
		return "", YtCfg{}, fmt.Errorf("unable to unmarshal ytcfg %s: %w", string(ytCfg), err)
	}

	var _initialData InitialData
	err = json.Unmarshal(initialData, &_initialData)
	if err != nil {
		return "", YtCfg{}, fmt.Errorf("unable to unmarshal initial data %s: %w", string(initialData), err)
	}

	continuations := _initialData.Contents.TwoColumnWatchNextResults.ConversationBar.LiveChatRenderer.Continuations
	if len(continuations) == 0 {
		return "", YtCfg{}, ErrStreamNotLive
	}
	reloadContinuationData := continuations[0].ReloadContinuationData
	initialContinuationInfo := reloadContinuationData.Continuation
	return initialContinuationInfo, _ytCfg, nil
}

func FetchContinuationChat(continuation string, ytCfg YtCfg) ([]ChatMessage, string, time.Duration, error) {
	chatMessages, continuation, timeoutMs, error := fetchChatMessages(continuation, ytCfg)
	if error != nil {
		return nil, "", time.Second, error
	}
	// Sleep for timeoutMs milliseconds sent by the server
	timeoutDuration := time.Second
	if timeoutMs > 0 {
		timeoutDuration = time.Duration(timeoutMs) * time.Millisecond
	}
	return chatMessages, continuation, timeoutDuration, nil
}

func AddCookies(cookies []*http.Cookie) {
	customCookies = cookies
}
