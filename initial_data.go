package YtChat

type InitialData struct {
	ResponseContext struct {
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MainAppWebResponseContext struct {
			LoggedOut     bool   `json:"loggedOut"`
			TrackingParam string `json:"trackingParam"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			YtConfigData struct {
				VisitorData           string `json:"visitorData"`
				RootVisualElementType int    `json:"rootVisualElementType"`
			} `json:"ytConfigData"`
			WebPrefetchData struct {
				NavigationEndpoints []struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							URL         string `json:"url"`
							WebPageType string `json:"webPageType"`
							RootVe      int    `json:"rootVe"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					WatchEndpoint struct {
						VideoID                              string `json:"videoId"`
						Params                               string `json:"params"`
						PlayerParams                         string `json:"playerParams"`
						WatchEndpointSupportedPrefetchConfig struct {
							PrefetchHintConfig struct {
								PrefetchPriority                            int `json:"prefetchPriority"`
								CountdownUIRelativeSecondsPrefetchCondition int `json:"countdownUiRelativeSecondsPrefetchCondition"`
							} `json:"prefetchHintConfig"`
						} `json:"watchEndpointSupportedPrefetchConfig"`
					} `json:"watchEndpoint"`
				} `json:"navigationEndpoints"`
			} `json:"webPrefetchData"`
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	Contents struct {
		TwoColumnWatchNextResults struct {
			Results struct {
				Results struct {
					Contents []struct {
						VideoPrimaryInfoRenderer struct {
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"title"`
							ViewCount struct {
								VideoViewCountRenderer struct {
									ViewCount struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"viewCount"`
									IsLive            bool   `json:"isLive"`
									OriginalViewCount string `json:"originalViewCount"`
								} `json:"videoViewCountRenderer"`
							} `json:"viewCount"`
							VideoActions struct {
								MenuRenderer struct {
									Items []struct {
										MenuServiceItemRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											ServiceEndpoint struct {
												ClickTrackingParams         string `json:"clickTrackingParams"`
												ShowEngagementPanelEndpoint struct {
													Identifier struct {
														Tag string `json:"tag"`
													} `json:"identifier"`
													GlobalConfiguration struct {
														Params string `json:"params"`
													} `json:"globalConfiguration"`
													EngagementPanelPresentationConfigs struct {
														EngagementPanelPopupPresentationConfig struct {
															PopupType string `json:"popupType"`
														} `json:"engagementPanelPopupPresentationConfig"`
													} `json:"engagementPanelPresentationConfigs"`
												} `json:"showEngagementPanelEndpoint"`
											} `json:"serviceEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"menuServiceItemRenderer"`
									} `json:"items"`
									TrackingParams  string `json:"trackingParams"`
									TopLevelButtons []struct {
										SegmentedLikeDislikeButtonViewModel struct {
											LikeButtonViewModel struct {
												LikeButtonViewModel struct {
													ToggleButtonViewModel struct {
														ToggleButtonViewModel struct {
															DefaultButtonViewModel struct {
																ButtonViewModel struct {
																	IconName string `json:"iconName"`
																	Title    string `json:"title"`
																	OnTap    struct {
																		SerialCommand struct {
																			Commands []struct {
																				LogGestureCommand struct {
																					GestureType    string `json:"gestureType"`
																					TrackingParams string `json:"trackingParams"`
																				} `json:"logGestureCommand,omitempty"`
																				InnertubeCommand struct {
																					ClickTrackingParams string `json:"clickTrackingParams"`
																					CommandMetadata     struct {
																						WebCommandMetadata struct {
																							IgnoreNavigation bool `json:"ignoreNavigation"`
																						} `json:"webCommandMetadata"`
																					} `json:"commandMetadata"`
																					ModalEndpoint struct {
																						Modal struct {
																							ModalWithTitleAndButtonRenderer struct {
																								Title struct {
																									SimpleText string `json:"simpleText"`
																								} `json:"title"`
																								Content struct {
																									SimpleText string `json:"simpleText"`
																								} `json:"content"`
																								Button struct {
																									ButtonRenderer struct {
																										Style      string `json:"style"`
																										Size       string `json:"size"`
																										IsDisabled bool   `json:"isDisabled"`
																										Text       struct {
																											SimpleText string `json:"simpleText"`
																										} `json:"text"`
																										NavigationEndpoint struct {
																											ClickTrackingParams string `json:"clickTrackingParams"`
																											CommandMetadata     struct {
																												WebCommandMetadata struct {
																													URL         string `json:"url"`
																													WebPageType string `json:"webPageType"`
																													RootVe      int    `json:"rootVe"`
																												} `json:"webCommandMetadata"`
																											} `json:"commandMetadata"`
																											SignInEndpoint struct {
																												NextEndpoint struct {
																													ClickTrackingParams string `json:"clickTrackingParams"`
																													CommandMetadata     struct {
																														WebCommandMetadata struct {
																															SendPost bool   `json:"sendPost"`
																															APIURL   string `json:"apiUrl"`
																														} `json:"webCommandMetadata"`
																													} `json:"commandMetadata"`
																													LikeEndpoint struct {
																														Status string `json:"status"`
																														Target struct {
																															VideoID string `json:"videoId"`
																														} `json:"target"`
																														LikeParams string `json:"likeParams"`
																													} `json:"likeEndpoint"`
																												} `json:"nextEndpoint"`
																												IdamTag string `json:"idamTag"`
																											} `json:"signInEndpoint"`
																										} `json:"navigationEndpoint"`
																										TrackingParams string `json:"trackingParams"`
																									} `json:"buttonRenderer"`
																								} `json:"button"`
																							} `json:"modalWithTitleAndButtonRenderer"`
																						} `json:"modal"`
																					} `json:"modalEndpoint"`
																				} `json:"innertubeCommand,omitempty"`
																			} `json:"commands"`
																		} `json:"serialCommand"`
																	} `json:"onTap"`
																	AccessibilityText string `json:"accessibilityText"`
																	Style             string `json:"style"`
																	TrackingParams    string `json:"trackingParams"`
																	IsFullWidth       bool   `json:"isFullWidth"`
																	Type              string `json:"type"`
																	ButtonSize        string `json:"buttonSize"`
																	AccessibilityID   string `json:"accessibilityId"`
																	Tooltip           string `json:"tooltip"`
																} `json:"buttonViewModel"`
															} `json:"defaultButtonViewModel"`
															ToggledButtonViewModel struct {
																ButtonViewModel struct {
																	IconName string `json:"iconName"`
																	Title    string `json:"title"`
																	OnTap    struct {
																		SerialCommand struct {
																			Commands []struct {
																				LogGestureCommand struct {
																					GestureType    string `json:"gestureType"`
																					TrackingParams string `json:"trackingParams"`
																				} `json:"logGestureCommand,omitempty"`
																				InnertubeCommand struct {
																					ClickTrackingParams string `json:"clickTrackingParams"`
																					CommandMetadata     struct {
																						WebCommandMetadata struct {
																							SendPost bool   `json:"sendPost"`
																							APIURL   string `json:"apiUrl"`
																						} `json:"webCommandMetadata"`
																					} `json:"commandMetadata"`
																					LikeEndpoint struct {
																						Status string `json:"status"`
																						Target struct {
																							VideoID string `json:"videoId"`
																						} `json:"target"`
																						RemoveLikeParams string `json:"removeLikeParams"`
																					} `json:"likeEndpoint"`
																				} `json:"innertubeCommand,omitempty"`
																			} `json:"commands"`
																		} `json:"serialCommand"`
																	} `json:"onTap"`
																	AccessibilityText string `json:"accessibilityText"`
																	Style             string `json:"style"`
																	TrackingParams    string `json:"trackingParams"`
																	IsFullWidth       bool   `json:"isFullWidth"`
																	Type              string `json:"type"`
																	ButtonSize        string `json:"buttonSize"`
																	AccessibilityID   string `json:"accessibilityId"`
																	Tooltip           string `json:"tooltip"`
																} `json:"buttonViewModel"`
															} `json:"toggledButtonViewModel"`
															Identifier         string `json:"identifier"`
															TrackingParams     string `json:"trackingParams"`
															IsTogglingDisabled bool   `json:"isTogglingDisabled"`
														} `json:"toggleButtonViewModel"`
													} `json:"toggleButtonViewModel"`
													LikeStatusEntityKey string `json:"likeStatusEntityKey"`
													LikeStatusEntity    struct {
														Key        string `json:"key"`
														LikeStatus string `json:"likeStatus"`
													} `json:"likeStatusEntity"`
												} `json:"likeButtonViewModel"`
											} `json:"likeButtonViewModel"`
											DislikeButtonViewModel struct {
												DislikeButtonViewModel struct {
													ToggleButtonViewModel struct {
														ToggleButtonViewModel struct {
															DefaultButtonViewModel struct {
																ButtonViewModel struct {
																	IconName string `json:"iconName"`
																	Title    string `json:"title"`
																	OnTap    struct {
																		SerialCommand struct {
																			Commands []struct {
																				LogGestureCommand struct {
																					GestureType    string `json:"gestureType"`
																					TrackingParams string `json:"trackingParams"`
																				} `json:"logGestureCommand,omitempty"`
																				InnertubeCommand struct {
																					ClickTrackingParams string `json:"clickTrackingParams"`
																					CommandMetadata     struct {
																						WebCommandMetadata struct {
																							IgnoreNavigation bool `json:"ignoreNavigation"`
																						} `json:"webCommandMetadata"`
																					} `json:"commandMetadata"`
																					ModalEndpoint struct {
																						Modal struct {
																							ModalWithTitleAndButtonRenderer struct {
																								Title struct {
																									SimpleText string `json:"simpleText"`
																								} `json:"title"`
																								Content struct {
																									SimpleText string `json:"simpleText"`
																								} `json:"content"`
																								Button struct {
																									ButtonRenderer struct {
																										Style      string `json:"style"`
																										Size       string `json:"size"`
																										IsDisabled bool   `json:"isDisabled"`
																										Text       struct {
																											SimpleText string `json:"simpleText"`
																										} `json:"text"`
																										NavigationEndpoint struct {
																											ClickTrackingParams string `json:"clickTrackingParams"`
																											CommandMetadata     struct {
																												WebCommandMetadata struct {
																													URL         string `json:"url"`
																													WebPageType string `json:"webPageType"`
																													RootVe      int    `json:"rootVe"`
																												} `json:"webCommandMetadata"`
																											} `json:"commandMetadata"`
																											SignInEndpoint struct {
																												NextEndpoint struct {
																													ClickTrackingParams string `json:"clickTrackingParams"`
																													CommandMetadata     struct {
																														WebCommandMetadata struct {
																															SendPost bool   `json:"sendPost"`
																															APIURL   string `json:"apiUrl"`
																														} `json:"webCommandMetadata"`
																													} `json:"commandMetadata"`
																													LikeEndpoint struct {
																														Status string `json:"status"`
																														Target struct {
																															VideoID string `json:"videoId"`
																														} `json:"target"`
																														DislikeParams string `json:"dislikeParams"`
																													} `json:"likeEndpoint"`
																												} `json:"nextEndpoint"`
																												IdamTag string `json:"idamTag"`
																											} `json:"signInEndpoint"`
																										} `json:"navigationEndpoint"`
																										TrackingParams string `json:"trackingParams"`
																									} `json:"buttonRenderer"`
																								} `json:"button"`
																							} `json:"modalWithTitleAndButtonRenderer"`
																						} `json:"modal"`
																					} `json:"modalEndpoint"`
																				} `json:"innertubeCommand,omitempty"`
																			} `json:"commands"`
																		} `json:"serialCommand"`
																	} `json:"onTap"`
																	AccessibilityText string `json:"accessibilityText"`
																	Style             string `json:"style"`
																	TrackingParams    string `json:"trackingParams"`
																	IsFullWidth       bool   `json:"isFullWidth"`
																	Type              string `json:"type"`
																	ButtonSize        string `json:"buttonSize"`
																	AccessibilityID   string `json:"accessibilityId"`
																	Tooltip           string `json:"tooltip"`
																} `json:"buttonViewModel"`
															} `json:"defaultButtonViewModel"`
															ToggledButtonViewModel struct {
																ButtonViewModel struct {
																	IconName string `json:"iconName"`
																	Title    string `json:"title"`
																	OnTap    struct {
																		SerialCommand struct {
																			Commands []struct {
																				LogGestureCommand struct {
																					GestureType    string `json:"gestureType"`
																					TrackingParams string `json:"trackingParams"`
																				} `json:"logGestureCommand,omitempty"`
																				InnertubeCommand struct {
																					ClickTrackingParams string `json:"clickTrackingParams"`
																					CommandMetadata     struct {
																						WebCommandMetadata struct {
																							SendPost bool   `json:"sendPost"`
																							APIURL   string `json:"apiUrl"`
																						} `json:"webCommandMetadata"`
																					} `json:"commandMetadata"`
																					LikeEndpoint struct {
																						Status string `json:"status"`
																						Target struct {
																							VideoID string `json:"videoId"`
																						} `json:"target"`
																						RemoveLikeParams string `json:"removeLikeParams"`
																					} `json:"likeEndpoint"`
																				} `json:"innertubeCommand,omitempty"`
																			} `json:"commands"`
																		} `json:"serialCommand"`
																	} `json:"onTap"`
																	AccessibilityText string `json:"accessibilityText"`
																	Style             string `json:"style"`
																	TrackingParams    string `json:"trackingParams"`
																	IsFullWidth       bool   `json:"isFullWidth"`
																	Type              string `json:"type"`
																	ButtonSize        string `json:"buttonSize"`
																	AccessibilityID   string `json:"accessibilityId"`
																	Tooltip           string `json:"tooltip"`
																} `json:"buttonViewModel"`
															} `json:"toggledButtonViewModel"`
															TrackingParams     string `json:"trackingParams"`
															IsTogglingDisabled bool   `json:"isTogglingDisabled"`
														} `json:"toggleButtonViewModel"`
													} `json:"toggleButtonViewModel"`
													DislikeEntityKey string `json:"dislikeEntityKey"`
												} `json:"dislikeButtonViewModel"`
											} `json:"dislikeButtonViewModel"`
											IconType        string `json:"iconType"`
											LikeCountEntity struct {
												Key              string `json:"key"`
												LikeCountIfLiked struct {
													Content string `json:"content"`
												} `json:"likeCountIfLiked"`
												LikeCountIfDisliked struct {
													Content string `json:"content"`
												} `json:"likeCountIfDisliked"`
												LikeCountIfIndifferent struct {
													Content string `json:"content"`
												} `json:"likeCountIfIndifferent"`
												ExpandedLikeCountIfLiked struct {
													Content string `json:"content"`
												} `json:"expandedLikeCountIfLiked"`
												ExpandedLikeCountIfDisliked struct {
													Content string `json:"content"`
												} `json:"expandedLikeCountIfDisliked"`
												ExpandedLikeCountIfIndifferent struct {
													Content string `json:"content"`
												} `json:"expandedLikeCountIfIndifferent"`
												LikeCountLabel struct {
													Content string `json:"content"`
												} `json:"likeCountLabel"`
												LikeButtonA11YText struct {
													Content string `json:"content"`
												} `json:"likeButtonA11yText"`
												LikeCountIfLikedNumber          string `json:"likeCountIfLikedNumber"`
												LikeCountIfDislikedNumber       string `json:"likeCountIfDislikedNumber"`
												LikeCountIfIndifferentNumber    string `json:"likeCountIfIndifferentNumber"`
												ShouldExpandLikeCount           bool   `json:"shouldExpandLikeCount"`
												SentimentFactoidA11YTextIfLiked struct {
													Content string `json:"content"`
												} `json:"sentimentFactoidA11yTextIfLiked"`
												SentimentFactoidA11YTextIfDisliked struct {
													Content string `json:"content"`
												} `json:"sentimentFactoidA11yTextIfDisliked"`
											} `json:"likeCountEntity"`
											DynamicLikeCountUpdateData struct {
												UpdateStatusKey               string `json:"updateStatusKey"`
												PlaceholderLikeCountValuesKey string `json:"placeholderLikeCountValuesKey"`
												UpdateDelayLoopID             string `json:"updateDelayLoopId"`
												UpdateDelaySec                int    `json:"updateDelaySec"`
											} `json:"dynamicLikeCountUpdateData"`
											TeasersOrderEntityKey string `json:"teasersOrderEntityKey"`
										} `json:"segmentedLikeDislikeButtonViewModel,omitempty"`
										ButtonViewModel struct {
											IconName string `json:"iconName"`
											Title    string `json:"title"`
											OnTap    struct {
												SerialCommand struct {
													Commands []struct {
														LogGestureCommand struct {
															GestureType    string `json:"gestureType"`
															TrackingParams string `json:"trackingParams"`
														} `json:"logGestureCommand,omitempty"`
														InnertubeCommand struct {
															ClickTrackingParams string `json:"clickTrackingParams"`
															CommandMetadata     struct {
																WebCommandMetadata struct {
																	SendPost bool   `json:"sendPost"`
																	APIURL   string `json:"apiUrl"`
																} `json:"webCommandMetadata"`
															} `json:"commandMetadata"`
															ShareEntityServiceEndpoint struct {
																SerializedShareEntity string `json:"serializedShareEntity"`
																Commands              []struct {
																	ClickTrackingParams string `json:"clickTrackingParams"`
																	OpenPopupAction     struct {
																		Popup struct {
																			UnifiedSharePanelRenderer struct {
																				TrackingParams     string `json:"trackingParams"`
																				ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
																			} `json:"unifiedSharePanelRenderer"`
																		} `json:"popup"`
																		PopupType string `json:"popupType"`
																		BeReused  bool   `json:"beReused"`
																	} `json:"openPopupAction"`
																} `json:"commands"`
															} `json:"shareEntityServiceEndpoint"`
														} `json:"innertubeCommand,omitempty"`
													} `json:"commands"`
												} `json:"serialCommand"`
											} `json:"onTap"`
											AccessibilityText string `json:"accessibilityText"`
											Style             string `json:"style"`
											TrackingParams    string `json:"trackingParams"`
											IsFullWidth       bool   `json:"isFullWidth"`
											Type              string `json:"type"`
											ButtonSize        string `json:"buttonSize"`
											State             string `json:"state"`
											AccessibilityID   string `json:"accessibilityId"`
											Tooltip           string `json:"tooltip"`
										} `json:"buttonViewModel,omitempty"`
									} `json:"topLevelButtons"`
									Accessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
									FlexibleItems []struct {
										MenuFlexibleItemRenderer struct {
											MenuItem struct {
												MenuServiceItemRenderer struct {
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													Icon struct {
														IconType string `json:"iconType"`
													} `json:"icon"`
													ServiceEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																IgnoreNavigation bool `json:"ignoreNavigation"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														ModalEndpoint struct {
															Modal struct {
																ModalWithTitleAndButtonRenderer struct {
																	Title struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"title"`
																	Content struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"content"`
																	Button struct {
																		ButtonRenderer struct {
																			Style      string `json:"style"`
																			Size       string `json:"size"`
																			IsDisabled bool   `json:"isDisabled"`
																			Text       struct {
																				SimpleText string `json:"simpleText"`
																			} `json:"text"`
																			NavigationEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						URL         string `json:"url"`
																						WebPageType string `json:"webPageType"`
																						RootVe      int    `json:"rootVe"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				SignInEndpoint struct {
																					NextEndpoint struct {
																						ClickTrackingParams string `json:"clickTrackingParams"`
																						CommandMetadata     struct {
																							WebCommandMetadata struct {
																								URL         string `json:"url"`
																								WebPageType string `json:"webPageType"`
																								RootVe      int    `json:"rootVe"`
																							} `json:"webCommandMetadata"`
																						} `json:"commandMetadata"`
																						WatchEndpoint struct {
																							VideoID                            string `json:"videoId"`
																							WatchEndpointSupportedOnesieConfig struct {
																								HTML5PlaybackOnesieConfig struct {
																									CommonConfig struct {
																										URL string `json:"url"`
																									} `json:"commonConfig"`
																								} `json:"html5PlaybackOnesieConfig"`
																							} `json:"watchEndpointSupportedOnesieConfig"`
																						} `json:"watchEndpoint"`
																					} `json:"nextEndpoint"`
																					IdamTag string `json:"idamTag"`
																				} `json:"signInEndpoint"`
																			} `json:"navigationEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"buttonRenderer"`
																	} `json:"button"`
																} `json:"modalWithTitleAndButtonRenderer"`
															} `json:"modal"`
														} `json:"modalEndpoint"`
													} `json:"serviceEndpoint"`
													TrackingParams string `json:"trackingParams"`
												} `json:"menuServiceItemRenderer"`
											} `json:"menuItem"`
											TopLevelButton struct {
												ButtonViewModel struct {
													IconName string `json:"iconName"`
													Title    string `json:"title"`
													OnTap    struct {
														SerialCommand struct {
															Commands []struct {
																LogGestureCommand struct {
																	GestureType    string `json:"gestureType"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"logGestureCommand,omitempty"`
																InnertubeCommand struct {
																	ClickTrackingParams string `json:"clickTrackingParams"`
																	CommandMetadata     struct {
																		WebCommandMetadata struct {
																			IgnoreNavigation bool `json:"ignoreNavigation"`
																		} `json:"webCommandMetadata"`
																	} `json:"commandMetadata"`
																	ModalEndpoint struct {
																		Modal struct {
																			ModalWithTitleAndButtonRenderer struct {
																				Title struct {
																					Runs []struct {
																						Text string `json:"text"`
																					} `json:"runs"`
																				} `json:"title"`
																				Content struct {
																					Runs []struct {
																						Text string `json:"text"`
																					} `json:"runs"`
																				} `json:"content"`
																				Button struct {
																					ButtonRenderer struct {
																						Style      string `json:"style"`
																						Size       string `json:"size"`
																						IsDisabled bool   `json:"isDisabled"`
																						Text       struct {
																							SimpleText string `json:"simpleText"`
																						} `json:"text"`
																						NavigationEndpoint struct {
																							ClickTrackingParams string `json:"clickTrackingParams"`
																							CommandMetadata     struct {
																								WebCommandMetadata struct {
																									URL         string `json:"url"`
																									WebPageType string `json:"webPageType"`
																									RootVe      int    `json:"rootVe"`
																								} `json:"webCommandMetadata"`
																							} `json:"commandMetadata"`
																							SignInEndpoint struct {
																								NextEndpoint struct {
																									ClickTrackingParams string `json:"clickTrackingParams"`
																									CommandMetadata     struct {
																										WebCommandMetadata struct {
																											URL         string `json:"url"`
																											WebPageType string `json:"webPageType"`
																											RootVe      int    `json:"rootVe"`
																										} `json:"webCommandMetadata"`
																									} `json:"commandMetadata"`
																									WatchEndpoint struct {
																										VideoID                            string `json:"videoId"`
																										WatchEndpointSupportedOnesieConfig struct {
																											HTML5PlaybackOnesieConfig struct {
																												CommonConfig struct {
																													URL string `json:"url"`
																												} `json:"commonConfig"`
																											} `json:"html5PlaybackOnesieConfig"`
																										} `json:"watchEndpointSupportedOnesieConfig"`
																									} `json:"watchEndpoint"`
																								} `json:"nextEndpoint"`
																								IdamTag string `json:"idamTag"`
																							} `json:"signInEndpoint"`
																						} `json:"navigationEndpoint"`
																						TrackingParams string `json:"trackingParams"`
																					} `json:"buttonRenderer"`
																				} `json:"button"`
																			} `json:"modalWithTitleAndButtonRenderer"`
																		} `json:"modal"`
																	} `json:"modalEndpoint"`
																} `json:"innertubeCommand,omitempty"`
															} `json:"commands"`
														} `json:"serialCommand"`
													} `json:"onTap"`
													AccessibilityText string `json:"accessibilityText"`
													Style             string `json:"style"`
													TrackingParams    string `json:"trackingParams"`
													IsFullWidth       bool   `json:"isFullWidth"`
													Type              string `json:"type"`
													ButtonSize        string `json:"buttonSize"`
													Tooltip           string `json:"tooltip"`
												} `json:"buttonViewModel"`
											} `json:"topLevelButton"`
										} `json:"menuFlexibleItemRenderer"`
									} `json:"flexibleItems"`
								} `json:"menuRenderer"`
							} `json:"videoActions"`
							TrackingParams          string `json:"trackingParams"`
							UpdatedMetadataEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										APIURL   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								UpdatedMetadataEndpoint struct {
									VideoID        string `json:"videoId"`
									InitialDelayMs int    `json:"initialDelayMs"`
									Params         string `json:"params"`
								} `json:"updatedMetadataEndpoint"`
							} `json:"updatedMetadataEndpoint"`
							DateText struct {
								SimpleText string `json:"simpleText"`
							} `json:"dateText"`
						} `json:"videoPrimaryInfoRenderer,omitempty"`
						VideoSecondaryInfoRenderer struct {
							Owner struct {
								VideoOwnerRenderer struct {
									Thumbnail struct {
										Thumbnails []struct {
											URL    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"thumbnails"`
									} `json:"thumbnail"`
									Title struct {
										Runs []struct {
											Text               string `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
														APIURL      string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												BrowseEndpoint struct {
													BrowseID         string `json:"browseId"`
													CanonicalBaseURL string `json:"canonicalBaseUrl"`
												} `json:"browseEndpoint"`
											} `json:"navigationEndpoint"`
										} `json:"runs"`
									} `json:"title"`
									SubscriptionButton struct {
										Type string `json:"type"`
									} `json:"subscriptionButton"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												APIURL      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseID         string `json:"browseId"`
											CanonicalBaseURL string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
									SubscriberCountText struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"subscriberCountText"`
									TrackingParams string `json:"trackingParams"`
									Badges         []struct {
										MetadataBadgeRenderer struct {
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Style             string `json:"style"`
											Tooltip           string `json:"tooltip"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"metadataBadgeRenderer"`
									} `json:"badges"`
									MembershipButton struct {
										TimedAnimationButtonRenderer struct {
											ButtonRenderer struct {
												ButtonRenderer struct {
													Style      string `json:"style"`
													Size       string `json:"size"`
													IsDisabled bool   `json:"isDisabled"`
													Text       struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													NavigationEndpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																IgnoreNavigation bool `json:"ignoreNavigation"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														ModalEndpoint struct {
															Modal struct {
																ModalWithTitleAndButtonRenderer struct {
																	Title struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"title"`
																	Content struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"content"`
																	Button struct {
																		ButtonRenderer struct {
																			Style      string `json:"style"`
																			Size       string `json:"size"`
																			IsDisabled bool   `json:"isDisabled"`
																			Text       struct {
																				SimpleText string `json:"simpleText"`
																			} `json:"text"`
																			NavigationEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						URL         string `json:"url"`
																						WebPageType string `json:"webPageType"`
																						RootVe      int    `json:"rootVe"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				SignInEndpoint struct {
																					Hack bool `json:"hack"`
																				} `json:"signInEndpoint"`
																			} `json:"navigationEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																		} `json:"buttonRenderer"`
																	} `json:"button"`
																} `json:"modalWithTitleAndButtonRenderer"`
															} `json:"modal"`
														} `json:"modalEndpoint"`
													} `json:"navigationEndpoint"`
													TrackingParams    string `json:"trackingParams"`
													AccessibilityData struct {
														AccessibilityData struct {
															Label string `json:"label"`
														} `json:"accessibilityData"`
													} `json:"accessibilityData"`
													TargetID string `json:"targetId"`
												} `json:"buttonRenderer"`
											} `json:"buttonRenderer"`
										} `json:"timedAnimationButtonRenderer"`
									} `json:"membershipButton"`
								} `json:"videoOwnerRenderer"`
							} `json:"owner"`
							SubscribeButton struct {
								SubscribeButtonRenderer struct {
									ButtonText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"buttonText"`
									Subscribed           bool   `json:"subscribed"`
									Enabled              bool   `json:"enabled"`
									Type                 string `json:"type"`
									ChannelID            string `json:"channelId"`
									ShowPreferences      bool   `json:"showPreferences"`
									SubscribedButtonText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"subscribedButtonText"`
									UnsubscribedButtonText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"unsubscribedButtonText"`
									TrackingParams        string `json:"trackingParams"`
									UnsubscribeButtonText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"unsubscribeButtonText"`
									SubscribeAccessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"subscribeAccessibility"`
									UnsubscribeAccessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"unsubscribeAccessibility"`
									NotificationPreferenceButton struct {
										SubscriptionNotificationToggleButtonRenderer struct {
											States []struct {
												StateID     int `json:"stateId"`
												NextStateID int `json:"nextStateId"`
												State       struct {
													ButtonRenderer struct {
														Style      string `json:"style"`
														Size       string `json:"size"`
														IsDisabled bool   `json:"isDisabled"`
														Icon       struct {
															IconType string `json:"iconType"`
														} `json:"icon"`
														Accessibility struct {
															Label string `json:"label"`
														} `json:"accessibility"`
														TrackingParams    string `json:"trackingParams"`
														AccessibilityData struct {
															AccessibilityData struct {
																Label string `json:"label"`
															} `json:"accessibilityData"`
														} `json:"accessibilityData"`
													} `json:"buttonRenderer"`
												} `json:"state"`
											} `json:"states"`
											CurrentStateID int    `json:"currentStateId"`
											TrackingParams string `json:"trackingParams"`
											Command        struct {
												ClickTrackingParams    string `json:"clickTrackingParams"`
												CommandExecutorCommand struct {
													Commands []struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														OpenPopupAction     struct {
															Popup struct {
																MenuPopupRenderer struct {
																	Items []struct {
																		MenuServiceItemRenderer struct {
																			Text struct {
																				SimpleText string `json:"simpleText"`
																			} `json:"text"`
																			Icon struct {
																				IconType string `json:"iconType"`
																			} `json:"icon"`
																			ServiceEndpoint struct {
																				ClickTrackingParams string `json:"clickTrackingParams"`
																				CommandMetadata     struct {
																					WebCommandMetadata struct {
																						SendPost bool   `json:"sendPost"`
																						APIURL   string `json:"apiUrl"`
																					} `json:"webCommandMetadata"`
																				} `json:"commandMetadata"`
																				ModifyChannelNotificationPreferenceEndpoint struct {
																					Params string `json:"params"`
																				} `json:"modifyChannelNotificationPreferenceEndpoint"`
																			} `json:"serviceEndpoint"`
																			TrackingParams string `json:"trackingParams"`
																			IsSelected     bool   `json:"isSelected"`
																		} `json:"menuServiceItemRenderer"`
																	} `json:"items"`
																} `json:"menuPopupRenderer"`
															} `json:"popup"`
															PopupType string `json:"popupType"`
														} `json:"openPopupAction"`
													} `json:"commands"`
												} `json:"commandExecutorCommand"`
											} `json:"command"`
											TargetID      string `json:"targetId"`
											SecondaryIcon struct {
												IconType string `json:"iconType"`
											} `json:"secondaryIcon"`
										} `json:"subscriptionNotificationToggleButtonRenderer"`
									} `json:"notificationPreferenceButton"`
									TargetID       string `json:"targetId"`
									SignInEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												IgnoreNavigation bool `json:"ignoreNavigation"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ModalEndpoint struct {
											Modal struct {
												ModalWithTitleAndButtonRenderer struct {
													Title struct {
														SimpleText string `json:"simpleText"`
													} `json:"title"`
													Content struct {
														SimpleText string `json:"simpleText"`
													} `json:"content"`
													Button struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Text       struct {
																SimpleText string `json:"simpleText"`
															} `json:"text"`
															NavigationEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		URL         string `json:"url"`
																		WebPageType string `json:"webPageType"`
																		RootVe      int    `json:"rootVe"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																SignInEndpoint struct {
																	NextEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				URL         string `json:"url"`
																				WebPageType string `json:"webPageType"`
																				RootVe      int    `json:"rootVe"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		WatchEndpoint struct {
																			VideoID                            string `json:"videoId"`
																			WatchEndpointSupportedOnesieConfig struct {
																				HTML5PlaybackOnesieConfig struct {
																					CommonConfig struct {
																						URL string `json:"url"`
																					} `json:"commonConfig"`
																				} `json:"html5PlaybackOnesieConfig"`
																			} `json:"watchEndpointSupportedOnesieConfig"`
																		} `json:"watchEndpoint"`
																	} `json:"nextEndpoint"`
																	ContinueAction string `json:"continueAction"`
																	IdamTag        string `json:"idamTag"`
																} `json:"signInEndpoint"`
															} `json:"navigationEndpoint"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"button"`
												} `json:"modalWithTitleAndButtonRenderer"`
											} `json:"modal"`
										} `json:"modalEndpoint"`
									} `json:"signInEndpoint"`
									SubscribedEntityKey  string `json:"subscribedEntityKey"`
									OnSubscribeEndpoints []struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												APIURL   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										SubscribeEndpoint struct {
											ChannelIds []string `json:"channelIds"`
											Params     string   `json:"params"`
										} `json:"subscribeEndpoint"`
									} `json:"onSubscribeEndpoints"`
									OnUnsubscribeEndpoints []struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool `json:"sendPost"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										SignalServiceEndpoint struct {
											Signal  string `json:"signal"`
											Actions []struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												OpenPopupAction     struct {
													Popup struct {
														ConfirmDialogRenderer struct {
															TrackingParams string `json:"trackingParams"`
															DialogMessages []struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"dialogMessages"`
															ConfirmButton struct {
																ButtonRenderer struct {
																	Style      string `json:"style"`
																	Size       string `json:"size"`
																	IsDisabled bool   `json:"isDisabled"`
																	Text       struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																	ServiceEndpoint struct {
																		ClickTrackingParams string `json:"clickTrackingParams"`
																		CommandMetadata     struct {
																			WebCommandMetadata struct {
																				SendPost bool   `json:"sendPost"`
																				APIURL   string `json:"apiUrl"`
																			} `json:"webCommandMetadata"`
																		} `json:"commandMetadata"`
																		UnsubscribeEndpoint struct {
																			ChannelIds []string `json:"channelIds"`
																			Params     string   `json:"params"`
																		} `json:"unsubscribeEndpoint"`
																	} `json:"serviceEndpoint"`
																	Accessibility struct {
																		Label string `json:"label"`
																	} `json:"accessibility"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"buttonRenderer"`
															} `json:"confirmButton"`
															CancelButton struct {
																ButtonRenderer struct {
																	Style      string `json:"style"`
																	Size       string `json:"size"`
																	IsDisabled bool   `json:"isDisabled"`
																	Text       struct {
																		Runs []struct {
																			Text string `json:"text"`
																		} `json:"runs"`
																	} `json:"text"`
																	Accessibility struct {
																		Label string `json:"label"`
																	} `json:"accessibility"`
																	TrackingParams string `json:"trackingParams"`
																} `json:"buttonRenderer"`
															} `json:"cancelButton"`
															PrimaryIsCancel bool `json:"primaryIsCancel"`
														} `json:"confirmDialogRenderer"`
													} `json:"popup"`
													PopupType string `json:"popupType"`
												} `json:"openPopupAction"`
											} `json:"actions"`
										} `json:"signalServiceEndpoint"`
									} `json:"onUnsubscribeEndpoints"`
								} `json:"subscribeButtonRenderer"`
							} `json:"subscribeButton"`
							MetadataRowContainer struct {
								MetadataRowContainerRenderer struct {
									Rows []struct {
										RichMetadataRowRenderer struct {
											Contents []struct {
												RichMetadataRenderer struct {
													Style     string `json:"style"`
													Thumbnail struct {
														Thumbnails []struct {
															URL    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"thumbnails"`
													} `json:"thumbnail"`
													Title struct {
														SimpleText string `json:"simpleText"`
													} `json:"title"`
													Subtitle struct {
														SimpleText string `json:"simpleText"`
													} `json:"subtitle"`
													CallToAction struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"callToAction"`
													CallToActionIcon struct {
														IconType string `json:"iconType"`
													} `json:"callToActionIcon"`
													Endpoint struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														CommandMetadata     struct {
															WebCommandMetadata struct {
																URL         string `json:"url"`
																WebPageType string `json:"webPageType"`
																RootVe      int    `json:"rootVe"`
																APIURL      string `json:"apiUrl"`
															} `json:"webCommandMetadata"`
														} `json:"commandMetadata"`
														BrowseEndpoint struct {
															BrowseID string `json:"browseId"`
														} `json:"browseEndpoint"`
													} `json:"endpoint"`
													TrackingParams string `json:"trackingParams"`
												} `json:"richMetadataRenderer"`
											} `json:"contents"`
											TrackingParams string `json:"trackingParams"`
										} `json:"richMetadataRowRenderer"`
									} `json:"rows"`
									CollapsedItemCount int    `json:"collapsedItemCount"`
									TrackingParams     string `json:"trackingParams"`
								} `json:"metadataRowContainerRenderer"`
							} `json:"metadataRowContainer"`
							ShowMoreText struct {
								SimpleText string `json:"simpleText"`
							} `json:"showMoreText"`
							ShowLessText struct {
								SimpleText string `json:"simpleText"`
							} `json:"showLessText"`
							TrackingParams            string `json:"trackingParams"`
							DefaultExpanded           bool   `json:"defaultExpanded"`
							DescriptionCollapsedLines int    `json:"descriptionCollapsedLines"`
							ShowMoreCommand           struct {
								ClickTrackingParams    string `json:"clickTrackingParams"`
								CommandExecutorCommand struct {
									Commands []struct {
										ClickTrackingParams                   string `json:"clickTrackingParams"`
										ChangeEngagementPanelVisibilityAction struct {
											TargetID   string `json:"targetId"`
											Visibility string `json:"visibility"`
										} `json:"changeEngagementPanelVisibilityAction,omitempty"`
										ScrollToEngagementPanelCommand struct {
											TargetID string `json:"targetId"`
										} `json:"scrollToEngagementPanelCommand,omitempty"`
									} `json:"commands"`
								} `json:"commandExecutorCommand"`
							} `json:"showMoreCommand"`
							ShowLessCommand struct {
								ClickTrackingParams                   string `json:"clickTrackingParams"`
								ChangeEngagementPanelVisibilityAction struct {
									TargetID   string `json:"targetId"`
									Visibility string `json:"visibility"`
								} `json:"changeEngagementPanelVisibilityAction"`
							} `json:"showLessCommand"`
							AttributedDescription struct {
								Content     string `json:"content"`
								CommandRuns []struct {
									StartIndex int `json:"startIndex"`
									Length     int `json:"length"`
									OnTap      struct {
										InnertubeCommand struct {
											ClickTrackingParams string `json:"clickTrackingParams"`
											CommandMetadata     struct {
												WebCommandMetadata struct {
													URL         string `json:"url"`
													WebPageType string `json:"webPageType"`
													RootVe      int    `json:"rootVe"`
												} `json:"webCommandMetadata"`
											} `json:"commandMetadata"`
											URLEndpoint struct {
												URL      string `json:"url"`
												Nofollow bool   `json:"nofollow"`
											} `json:"urlEndpoint"`
										} `json:"innertubeCommand"`
									} `json:"onTap"`
									OnTapOptions struct {
										AccessibilityInfo struct {
											AccessibilityLabel string `json:"accessibilityLabel"`
										} `json:"accessibilityInfo"`
									} `json:"onTapOptions,omitempty"`
								} `json:"commandRuns"`
								StyleRuns []struct {
									StartIndex         int `json:"startIndex"`
									Length             int `json:"length"`
									StyleRunExtensions struct {
										StyleRunColorMapExtension struct {
											ColorMap []struct {
												Key   string `json:"key"`
												Value int64  `json:"value"`
											} `json:"colorMap"`
										} `json:"styleRunColorMapExtension"`
									} `json:"styleRunExtensions"`
									FontFamilyName string `json:"fontFamilyName,omitempty"`
								} `json:"styleRuns"`
								AttachmentRuns []struct {
									StartIndex int `json:"startIndex"`
									Length     int `json:"length"`
									Element    struct {
										Type struct {
											ImageType struct {
												Image struct {
													Sources []struct {
														URL string `json:"url"`
													} `json:"sources"`
												} `json:"image"`
											} `json:"imageType"`
										} `json:"type"`
										Properties struct {
											LayoutProperties struct {
												Height struct {
													Value int    `json:"value"`
													Unit  string `json:"unit"`
												} `json:"height"`
												Width struct {
													Value int    `json:"value"`
													Unit  string `json:"unit"`
												} `json:"width"`
												Margin struct {
													Top struct {
														Value float64 `json:"value"`
														Unit  string  `json:"unit"`
													} `json:"top"`
												} `json:"margin"`
											} `json:"layoutProperties"`
										} `json:"properties"`
									} `json:"element"`
									Alignment string `json:"alignment"`
								} `json:"attachmentRuns"`
								DecorationRuns []struct {
									TextDecorator struct {
										HighlightTextDecorator struct {
											StartIndex                       int `json:"startIndex"`
											Length                           int `json:"length"`
											BackgroundCornerRadius           int `json:"backgroundCornerRadius"`
											BottomPadding                    int `json:"bottomPadding"`
											HighlightTextDecoratorExtensions struct {
												HighlightTextDecoratorColorMapExtension struct {
													ColorMap []struct {
														Key   string `json:"key"`
														Value int    `json:"value"`
													} `json:"colorMap"`
												} `json:"highlightTextDecoratorColorMapExtension"`
											} `json:"highlightTextDecoratorExtensions"`
										} `json:"highlightTextDecorator"`
									} `json:"textDecorator"`
								} `json:"decorationRuns"`
							} `json:"attributedDescription"`
							HeaderRuns []struct {
								StartIndex    int    `json:"startIndex"`
								Length        int    `json:"length"`
								HeaderMapping string `json:"headerMapping"`
							} `json:"headerRuns"`
						} `json:"videoSecondaryInfoRenderer,omitempty"`
						CompositeVideoPrimaryInfoRenderer struct {
						} `json:"compositeVideoPrimaryInfoRenderer,omitempty"`
						ItemSectionRenderer struct {
							Contents []struct {
								VideoMetadataCarouselViewModel struct {
									CarouselTitles []struct {
										CarouselTitleViewModel struct {
											Title          string `json:"title"`
											PreviousButton struct {
												ButtonViewModel struct {
													IconName          string `json:"iconName"`
													AccessibilityText string `json:"accessibilityText"`
													Style             string `json:"style"`
													TrackingParams    string `json:"trackingParams"`
													Type              string `json:"type"`
													ButtonSize        string `json:"buttonSize"`
													State             string `json:"state"`
												} `json:"buttonViewModel"`
											} `json:"previousButton"`
											NextButton struct {
												ButtonViewModel struct {
													IconName          string `json:"iconName"`
													AccessibilityText string `json:"accessibilityText"`
													Style             string `json:"style"`
													TrackingParams    string `json:"trackingParams"`
													Type              string `json:"type"`
													ButtonSize        string `json:"buttonSize"`
													State             string `json:"state"`
												} `json:"buttonViewModel"`
											} `json:"nextButton"`
										} `json:"carouselTitleViewModel"`
									} `json:"carouselTitles"`
									CarouselItems []struct {
										CarouselItemViewModel struct {
											ItemType     string `json:"itemType"`
											CarouselItem struct {
												TextCarouselItemViewModel struct {
													IconName string `json:"iconName"`
													Text     struct {
														Content string `json:"content"`
													} `json:"text"`
													OnTap struct {
														InnertubeCommand struct {
															ClickTrackingParams             string `json:"clickTrackingParams"`
															SetLiveChatCollapsedStateAction struct {
															} `json:"setLiveChatCollapsedStateAction"`
														} `json:"innertubeCommand"`
													} `json:"onTap"`
													TrackingParams string `json:"trackingParams"`
													Button         struct {
														ButtonViewModel struct {
															Title string `json:"title"`
															OnTap struct {
																InnertubeCommand struct {
																	ClickTrackingParams             string `json:"clickTrackingParams"`
																	SetLiveChatCollapsedStateAction struct {
																	} `json:"setLiveChatCollapsedStateAction"`
																} `json:"innertubeCommand"`
															} `json:"onTap"`
															Style          string `json:"style"`
															TrackingParams string `json:"trackingParams"`
															Type           string `json:"type"`
															ButtonSize     string `json:"buttonSize"`
														} `json:"buttonViewModel"`
													} `json:"button"`
												} `json:"textCarouselItemViewModel"`
											} `json:"carouselItem"`
										} `json:"carouselItemViewModel"`
									} `json:"carouselItems"`
								} `json:"videoMetadataCarouselViewModel"`
							} `json:"contents"`
							TrackingParams string `json:"trackingParams"`
						} `json:"itemSectionRenderer,omitempty"`
					} `json:"contents"`
					TrackingParams string `json:"trackingParams"`
				} `json:"results"`
			} `json:"results"`
			SecondaryResults struct {
				SecondaryResults struct {
					Results []struct {
						LockupViewModel struct {
							ContentImage struct {
								ThumbnailViewModel struct {
									Image struct {
										Sources []struct {
											URL    string `json:"url"`
											Width  int    `json:"width"`
											Height int    `json:"height"`
										} `json:"sources"`
									} `json:"image"`
									Overlays []struct {
										ThumbnailOverlayBadgeViewModel struct {
											ThumbnailBadges []struct {
												ThumbnailBadgeViewModel struct {
													Text                         string `json:"text"`
													BadgeStyle                   string `json:"badgeStyle"`
													AnimationActivationTargetID  string `json:"animationActivationTargetId"`
													AnimationActivationEntityKey string `json:"animationActivationEntityKey"`
													LottieData                   struct {
														URL      string `json:"url"`
														Settings struct {
															Loop     bool `json:"loop"`
															Autoplay bool `json:"autoplay"`
														} `json:"settings"`
													} `json:"lottieData"`
													AnimatedText                          string `json:"animatedText"`
													AnimationActivationEntitySelectorType string `json:"animationActivationEntitySelectorType"`
													RendererContext                       struct {
														AccessibilityContext struct {
															Label string `json:"label"`
														} `json:"accessibilityContext"`
													} `json:"rendererContext"`
												} `json:"thumbnailBadgeViewModel"`
											} `json:"thumbnailBadges"`
											Position string `json:"position"`
										} `json:"thumbnailOverlayBadgeViewModel,omitempty"`
										ThumbnailHoverOverlayToggleActionsViewModel struct {
											Buttons []struct {
												ToggleButtonViewModel struct {
													DefaultButtonViewModel struct {
														ButtonViewModel struct {
															IconName string `json:"iconName"`
															OnTap    struct {
																InnertubeCommand struct {
																	ClickTrackingParams string `json:"clickTrackingParams"`
																	CommandMetadata     struct {
																		WebCommandMetadata struct {
																			SendPost bool   `json:"sendPost"`
																			APIURL   string `json:"apiUrl"`
																		} `json:"webCommandMetadata"`
																	} `json:"commandMetadata"`
																	PlaylistEditEndpoint struct {
																		PlaylistID string `json:"playlistId"`
																		Actions    []struct {
																			AddedVideoID string `json:"addedVideoId"`
																			Action       string `json:"action"`
																		} `json:"actions"`
																	} `json:"playlistEditEndpoint"`
																} `json:"innertubeCommand"`
															} `json:"onTap"`
															AccessibilityText string `json:"accessibilityText"`
															Style             string `json:"style"`
															TrackingParams    string `json:"trackingParams"`
															Type              string `json:"type"`
															ButtonSize        string `json:"buttonSize"`
															State             string `json:"state"`
														} `json:"buttonViewModel"`
													} `json:"defaultButtonViewModel"`
													ToggledButtonViewModel struct {
														ButtonViewModel struct {
															IconName string `json:"iconName"`
															OnTap    struct {
																InnertubeCommand struct {
																	ClickTrackingParams string `json:"clickTrackingParams"`
																	CommandMetadata     struct {
																		WebCommandMetadata struct {
																			SendPost bool   `json:"sendPost"`
																			APIURL   string `json:"apiUrl"`
																		} `json:"webCommandMetadata"`
																	} `json:"commandMetadata"`
																	PlaylistEditEndpoint struct {
																		PlaylistID string `json:"playlistId"`
																		Actions    []struct {
																			Action         string `json:"action"`
																			RemovedVideoID string `json:"removedVideoId"`
																		} `json:"actions"`
																	} `json:"playlistEditEndpoint"`
																} `json:"innertubeCommand"`
															} `json:"onTap"`
															AccessibilityText string `json:"accessibilityText"`
															Style             string `json:"style"`
															TrackingParams    string `json:"trackingParams"`
															Type              string `json:"type"`
															ButtonSize        string `json:"buttonSize"`
															State             string `json:"state"`
														} `json:"buttonViewModel"`
													} `json:"toggledButtonViewModel"`
													IsToggled      bool   `json:"isToggled"`
													TrackingParams string `json:"trackingParams"`
												} `json:"toggleButtonViewModel"`
											} `json:"buttons"`
										} `json:"thumbnailHoverOverlayToggleActionsViewModel,omitempty"`
									} `json:"overlays"`
								} `json:"thumbnailViewModel"`
							} `json:"contentImage"`
							Metadata struct {
								LockupMetadataViewModel struct {
									Title struct {
										Content string `json:"content"`
									} `json:"title"`
									Image struct {
										DecoratedAvatarViewModel struct {
											Avatar struct {
												AvatarViewModel struct {
													Image struct {
														Sources []struct {
															URL    string `json:"url"`
															Width  int    `json:"width"`
															Height int    `json:"height"`
														} `json:"sources"`
													} `json:"image"`
													AvatarImageSize string `json:"avatarImageSize"`
												} `json:"avatarViewModel"`
											} `json:"avatar"`
											A11YLabel       string `json:"a11yLabel"`
											RendererContext struct {
												CommandContext struct {
													OnTap struct {
														InnertubeCommand struct {
															ClickTrackingParams string `json:"clickTrackingParams"`
															CommandMetadata     struct {
																WebCommandMetadata struct {
																	URL         string `json:"url"`
																	WebPageType string `json:"webPageType"`
																	RootVe      int    `json:"rootVe"`
																	APIURL      string `json:"apiUrl"`
																} `json:"webCommandMetadata"`
															} `json:"commandMetadata"`
															BrowseEndpoint struct {
																BrowseID         string `json:"browseId"`
																CanonicalBaseURL string `json:"canonicalBaseUrl"`
															} `json:"browseEndpoint"`
														} `json:"innertubeCommand"`
													} `json:"onTap"`
												} `json:"commandContext"`
											} `json:"rendererContext"`
										} `json:"decoratedAvatarViewModel"`
									} `json:"image"`
									Metadata struct {
										ContentMetadataViewModel struct {
											MetadataRows []struct {
												MetadataParts []struct {
													Text struct {
														Content   string `json:"content"`
														StyleRuns []struct {
															StartIndex         int `json:"startIndex"`
															StyleRunExtensions struct {
																StyleRunColorMapExtension struct {
																	ColorMap []struct {
																		Key   string `json:"key"`
																		Value int64  `json:"value"`
																	} `json:"colorMap"`
																} `json:"styleRunColorMapExtension"`
															} `json:"styleRunExtensions"`
														} `json:"styleRuns"`
														AttachmentRuns []struct {
															StartIndex int `json:"startIndex"`
															Length     int `json:"length"`
															Element    struct {
																Type struct {
																	ImageType struct {
																		Image struct {
																			Sources []struct {
																				ClientResource struct {
																					ImageName string `json:"imageName"`
																				} `json:"clientResource"`
																				Width  int `json:"width"`
																				Height int `json:"height"`
																			} `json:"sources"`
																		} `json:"image"`
																	} `json:"imageType"`
																} `json:"type"`
																Properties struct {
																	LayoutProperties struct {
																		Height struct {
																			Value int    `json:"value"`
																			Unit  string `json:"unit"`
																		} `json:"height"`
																		Width struct {
																			Value int    `json:"value"`
																			Unit  string `json:"unit"`
																		} `json:"width"`
																		Margin struct {
																			Left struct {
																				Value int    `json:"value"`
																				Unit  string `json:"unit"`
																			} `json:"left"`
																		} `json:"margin"`
																	} `json:"layoutProperties"`
																} `json:"properties"`
															} `json:"element"`
															Alignment string `json:"alignment"`
														} `json:"attachmentRuns"`
													} `json:"text"`
												} `json:"metadataParts,omitempty"`
												Badges []struct {
													BadgeViewModel struct {
														BadgeText      string `json:"badgeText"`
														BadgeStyle     string `json:"badgeStyle"`
														TrackingParams string `json:"trackingParams"`
													} `json:"badgeViewModel"`
												} `json:"badges,omitempty"`
											} `json:"metadataRows"`
											Delimiter string `json:"delimiter"`
										} `json:"contentMetadataViewModel"`
									} `json:"metadata"`
									MenuButton struct {
										ButtonViewModel struct {
											IconName string `json:"iconName"`
											OnTap    struct {
												InnertubeCommand struct {
													ClickTrackingParams string `json:"clickTrackingParams"`
													ShowSheetCommand    struct {
														PanelLoadingStrategy struct {
															InlineContent struct {
																SheetViewModel struct {
																	Content struct {
																		ListViewModel struct {
																			ListItems []struct {
																				ListItemViewModel struct {
																					Title struct {
																						Content string `json:"content"`
																					} `json:"title"`
																					LeadingImage struct {
																						Sources []struct {
																							ClientResource struct {
																								ImageName string `json:"imageName"`
																							} `json:"clientResource"`
																						} `json:"sources"`
																					} `json:"leadingImage"`
																					RendererContext struct {
																						LoggingContext struct {
																							LoggingDirectives struct {
																								TrackingParams string `json:"trackingParams"`
																								Visibility     struct {
																									Types string `json:"types"`
																								} `json:"visibility"`
																							} `json:"loggingDirectives"`
																						} `json:"loggingContext"`
																						CommandContext struct {
																							OnTap struct {
																								InnertubeCommand struct {
																									ClickTrackingParams string `json:"clickTrackingParams"`
																									CommandMetadata     struct {
																										WebCommandMetadata struct {
																											SendPost bool `json:"sendPost"`
																										} `json:"webCommandMetadata"`
																									} `json:"commandMetadata"`
																									SignalServiceEndpoint struct {
																										Signal  string `json:"signal"`
																										Actions []struct {
																											ClickTrackingParams  string `json:"clickTrackingParams"`
																											AddToPlaylistCommand struct {
																												OpenMiniplayer      bool   `json:"openMiniplayer"`
																												VideoID             string `json:"videoId"`
																												ListType            string `json:"listType"`
																												OnCreateListCommand struct {
																													ClickTrackingParams string `json:"clickTrackingParams"`
																													CommandMetadata     struct {
																														WebCommandMetadata struct {
																															SendPost bool   `json:"sendPost"`
																															APIURL   string `json:"apiUrl"`
																														} `json:"webCommandMetadata"`
																													} `json:"commandMetadata"`
																													CreatePlaylistServiceEndpoint struct {
																														VideoIds []string `json:"videoIds"`
																														Params   string   `json:"params"`
																													} `json:"createPlaylistServiceEndpoint"`
																												} `json:"onCreateListCommand"`
																												VideoIds     []string `json:"videoIds"`
																												VideoCommand struct {
																													ClickTrackingParams string `json:"clickTrackingParams"`
																													CommandMetadata     struct {
																														WebCommandMetadata struct {
																															URL         string `json:"url"`
																															WebPageType string `json:"webPageType"`
																															RootVe      int    `json:"rootVe"`
																														} `json:"webCommandMetadata"`
																													} `json:"commandMetadata"`
																													WatchEndpoint struct {
																														VideoID                            string `json:"videoId"`
																														WatchEndpointSupportedOnesieConfig struct {
																															HTML5PlaybackOnesieConfig struct {
																																CommonConfig struct {
																																	URL string `json:"url"`
																																} `json:"commonConfig"`
																															} `json:"html5PlaybackOnesieConfig"`
																														} `json:"watchEndpointSupportedOnesieConfig"`
																													} `json:"watchEndpoint"`
																												} `json:"videoCommand"`
																											} `json:"addToPlaylistCommand"`
																										} `json:"actions"`
																									} `json:"signalServiceEndpoint"`
																								} `json:"innertubeCommand"`
																							} `json:"onTap"`
																						} `json:"commandContext"`
																					} `json:"rendererContext"`
																				} `json:"listItemViewModel"`
																			} `json:"listItems"`
																		} `json:"listViewModel"`
																	} `json:"content"`
																} `json:"sheetViewModel"`
															} `json:"inlineContent"`
														} `json:"panelLoadingStrategy"`
													} `json:"showSheetCommand"`
												} `json:"innertubeCommand"`
											} `json:"onTap"`
											AccessibilityText string `json:"accessibilityText"`
											Style             string `json:"style"`
											TrackingParams    string `json:"trackingParams"`
											Type              string `json:"type"`
											ButtonSize        string `json:"buttonSize"`
											State             string `json:"state"`
										} `json:"buttonViewModel"`
									} `json:"menuButton"`
								} `json:"lockupMetadataViewModel"`
							} `json:"metadata"`
							ContentID       string `json:"contentId"`
							ContentType     string `json:"contentType"`
							RendererContext struct {
								LoggingContext struct {
									LoggingDirectives struct {
										TrackingParams string `json:"trackingParams"`
										Visibility     struct {
											Types string `json:"types"`
										} `json:"visibility"`
									} `json:"loggingDirectives"`
								} `json:"loggingContext"`
								AccessibilityContext struct {
									Label string `json:"label"`
								} `json:"accessibilityContext"`
								CommandContext struct {
									OnTap struct {
										InnertubeCommand struct {
											ClickTrackingParams string `json:"clickTrackingParams"`
											CommandMetadata     struct {
												WebCommandMetadata struct {
													URL         string `json:"url"`
													WebPageType string `json:"webPageType"`
													RootVe      int    `json:"rootVe"`
												} `json:"webCommandMetadata"`
											} `json:"commandMetadata"`
											WatchEndpoint struct {
												VideoID                            string `json:"videoId"`
												Nofollow                           bool   `json:"nofollow"`
												WatchEndpointSupportedOnesieConfig struct {
													HTML5PlaybackOnesieConfig struct {
														CommonConfig struct {
															URL string `json:"url"`
														} `json:"commonConfig"`
													} `json:"html5PlaybackOnesieConfig"`
												} `json:"watchEndpointSupportedOnesieConfig"`
											} `json:"watchEndpoint"`
										} `json:"innertubeCommand"`
									} `json:"onTap"`
								} `json:"commandContext"`
							} `json:"rendererContext"`
						} `json:"lockupViewModel,omitempty"`
						ContinuationItemRenderer struct {
							Trigger              string `json:"trigger"`
							ContinuationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										APIURL   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								ContinuationCommand struct {
									Token   string `json:"token"`
									Request string `json:"request"`
								} `json:"continuationCommand"`
							} `json:"continuationEndpoint"`
							Button struct {
								ButtonRenderer struct {
									Style      string `json:"style"`
									Size       string `json:"size"`
									IsDisabled bool   `json:"isDisabled"`
									Text       struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"text"`
									TrackingParams string `json:"trackingParams"`
									Command        struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool   `json:"sendPost"`
												APIURL   string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										ContinuationCommand struct {
											Token   string `json:"token"`
											Request string `json:"request"`
										} `json:"continuationCommand"`
									} `json:"command"`
								} `json:"buttonRenderer"`
							} `json:"button"`
						} `json:"continuationItemRenderer,omitempty"`
					} `json:"results"`
					TrackingParams string `json:"trackingParams"`
					TargetID       string `json:"targetId"`
				} `json:"secondaryResults"`
			} `json:"secondaryResults"`
			Autoplay struct {
				Autoplay struct {
					Sets []struct {
						Mode          string `json:"mode"`
						AutoplayVideo struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									URL         string `json:"url"`
									WebPageType string `json:"webPageType"`
									RootVe      int    `json:"rootVe"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							WatchEndpoint struct {
								VideoID                              string `json:"videoId"`
								Params                               string `json:"params"`
								PlayerParams                         string `json:"playerParams"`
								WatchEndpointSupportedPrefetchConfig struct {
									PrefetchHintConfig struct {
										PrefetchPriority                            int `json:"prefetchPriority"`
										CountdownUIRelativeSecondsPrefetchCondition int `json:"countdownUiRelativeSecondsPrefetchCondition"`
									} `json:"prefetchHintConfig"`
								} `json:"watchEndpointSupportedPrefetchConfig"`
							} `json:"watchEndpoint"`
						} `json:"autoplayVideo"`
					} `json:"sets"`
					CountDownSecs  int    `json:"countDownSecs"`
					TrackingParams string `json:"trackingParams"`
				} `json:"autoplay"`
			} `json:"autoplay"`
			ConversationBar struct {
				LiveChatRenderer struct {
					Continuations []struct {
						ReloadContinuationData struct {
							Continuation        string `json:"continuation"`
							ClickTrackingParams string `json:"clickTrackingParams"`
						} `json:"reloadContinuationData"`
					} `json:"continuations"`
					Header struct {
						LiveChatHeaderRenderer struct {
							OverflowMenu struct {
								MenuRenderer struct {
									Items []struct {
										MenuServiceItemRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											ServiceEndpoint struct {
												ClickTrackingParams              string `json:"clickTrackingParams"`
												ShowLiveChatParticipantsEndpoint struct {
													Hack bool `json:"hack"`
												} `json:"showLiveChatParticipantsEndpoint"`
											} `json:"serviceEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"menuServiceItemRenderer,omitempty"`
										ClientSideToggleMenuItemRenderer struct {
											DefaultText struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"defaultText"`
											DefaultIcon struct {
												IconType string `json:"iconType"`
											} `json:"defaultIcon"`
											ToggledText struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"toggledText"`
											ToggledIcon struct {
												IconType string `json:"iconType"`
											} `json:"toggledIcon"`
											MenuItemIdentifier string `json:"menuItemIdentifier"`
											Command            struct {
												ClickTrackingParams              string `json:"clickTrackingParams"`
												ToggleLiveChatTimestampsEndpoint struct {
													Hack bool `json:"hack"`
												} `json:"toggleLiveChatTimestampsEndpoint"`
											} `json:"command"`
										} `json:"clientSideToggleMenuItemRenderer,omitempty"`
										MenuNavigationItemRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														IgnoreNavigation bool `json:"ignoreNavigation"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												UserFeedbackEndpoint struct {
													Hack             bool   `json:"hack"`
													BucketIdentifier string `json:"bucketIdentifier"`
												} `json:"userFeedbackEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"menuNavigationItemRenderer,omitempty"`
									} `json:"items"`
									TrackingParams string `json:"trackingParams"`
									Accessibility  struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
								} `json:"menuRenderer"`
							} `json:"overflowMenu"`
							CollapseButton struct {
								ButtonRenderer struct {
									Style      string `json:"style"`
									Size       string `json:"size"`
									IsDisabled bool   `json:"isDisabled"`
									Icon       struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									Accessibility struct {
										Label string `json:"label"`
									} `json:"accessibility"`
									TrackingParams string `json:"trackingParams"`
									Command        struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool `json:"sendPost"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										SignalServiceEndpoint struct {
											Signal  string `json:"signal"`
											Actions []struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												SignalAction        struct {
													Signal string `json:"signal"`
												} `json:"signalAction"`
											} `json:"actions"`
										} `json:"signalServiceEndpoint"`
									} `json:"command"`
								} `json:"buttonRenderer"`
							} `json:"collapseButton"`
							ViewSelector struct {
								SortFilterSubMenuRenderer struct {
									SubMenuItems []struct {
										Title        string `json:"title"`
										Selected     bool   `json:"selected"`
										Continuation struct {
											ReloadContinuationData struct {
												Continuation        string `json:"continuation"`
												ClickTrackingParams string `json:"clickTrackingParams"`
											} `json:"reloadContinuationData"`
										} `json:"continuation"`
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										Subtitle       string `json:"subtitle"`
										TrackingParams string `json:"trackingParams"`
									} `json:"subMenuItems"`
									Accessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
									TrackingParams string `json:"trackingParams"`
									TargetID       string `json:"targetId"`
								} `json:"sortFilterSubMenuRenderer"`
							} `json:"viewSelector"`
						} `json:"liveChatHeaderRenderer"`
					} `json:"header"`
					TrackingParams string `json:"trackingParams"`
					ClientMessages struct {
						ReconnectMessage struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"reconnectMessage"`
						UnableToReconnectMessage struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"unableToReconnectMessage"`
						FatalError struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"fatalError"`
						ReconnectedMessage struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"reconnectedMessage"`
						GenericError struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"genericError"`
					} `json:"clientMessages"`
					InitialDisplayState string `json:"initialDisplayState"`
					ShowButton          struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"showButton"`
				} `json:"liveChatRenderer"`
			} `json:"conversationBar"`
		} `json:"twoColumnWatchNextResults"`
	} `json:"contents"`
	CurrentVideoEndpoint struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
		CommandMetadata     struct {
			WebCommandMetadata struct {
				URL         string `json:"url"`
				WebPageType string `json:"webPageType"`
				RootVe      int    `json:"rootVe"`
			} `json:"webCommandMetadata"`
		} `json:"commandMetadata"`
		WatchEndpoint struct {
			VideoID                            string `json:"videoId"`
			WatchEndpointSupportedOnesieConfig struct {
				HTML5PlaybackOnesieConfig struct {
					CommonConfig struct {
						URL string `json:"url"`
					} `json:"commonConfig"`
				} `json:"html5PlaybackOnesieConfig"`
			} `json:"watchEndpointSupportedOnesieConfig"`
		} `json:"watchEndpoint"`
	} `json:"currentVideoEndpoint"`
	TrackingParams string `json:"trackingParams"`
	PlayerOverlays struct {
		PlayerOverlayRenderer struct {
			EndScreen struct {
				WatchNextEndScreenRenderer struct {
					Results []struct {
						EndScreenVideoRenderer struct {
							VideoID   string `json:"videoId"`
							Thumbnail struct {
								Thumbnails []struct {
									URL    string `json:"url"`
									Width  int    `json:"width"`
									Height int    `json:"height"`
								} `json:"thumbnails"`
							} `json:"thumbnail"`
							Title struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"title"`
							ShortBylineText struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
												APIURL      string `json:"apiUrl"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										BrowseEndpoint struct {
											BrowseID         string `json:"browseId"`
											CanonicalBaseURL string `json:"canonicalBaseUrl"`
										} `json:"browseEndpoint"`
									} `json:"navigationEndpoint"`
								} `json:"runs"`
							} `json:"shortBylineText"`
							LengthText struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"lengthText"`
							LengthInSeconds    int `json:"lengthInSeconds"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoID                            string `json:"videoId"`
									WatchEndpointSupportedOnesieConfig struct {
										HTML5PlaybackOnesieConfig struct {
											CommonConfig struct {
												URL string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
								} `json:"watchEndpoint"`
							} `json:"navigationEndpoint"`
							TrackingParams     string `json:"trackingParams"`
							ShortViewCountText struct {
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
								SimpleText string `json:"simpleText"`
							} `json:"shortViewCountText"`
							PublishedTimeText struct {
								SimpleText string `json:"simpleText"`
							} `json:"publishedTimeText"`
							ThumbnailOverlays []struct {
								ThumbnailOverlayTimeStatusRenderer struct {
									Text struct {
										Accessibility struct {
											AccessibilityData struct {
												Label string `json:"label"`
											} `json:"accessibilityData"`
										} `json:"accessibility"`
										SimpleText string `json:"simpleText"`
									} `json:"text"`
									Style string `json:"style"`
								} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
								ThumbnailOverlayNowPlayingRenderer struct {
									Text struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"text"`
								} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
							} `json:"thumbnailOverlays"`
						} `json:"endScreenVideoRenderer"`
					} `json:"results"`
					Title struct {
						SimpleText string `json:"simpleText"`
					} `json:"title"`
					TrackingParams string `json:"trackingParams"`
				} `json:"watchNextEndScreenRenderer"`
			} `json:"endScreen"`
			Autoplay struct {
				PlayerOverlayAutoplayRenderer struct {
					Title struct {
						SimpleText string `json:"simpleText"`
					} `json:"title"`
					VideoTitle struct {
						Accessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
						SimpleText string `json:"simpleText"`
					} `json:"videoTitle"`
					Byline struct {
						Runs []struct {
							Text               string `json:"text"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
										APIURL      string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								BrowseEndpoint struct {
									BrowseID         string `json:"browseId"`
									CanonicalBaseURL string `json:"canonicalBaseUrl"`
								} `json:"browseEndpoint"`
							} `json:"navigationEndpoint"`
						} `json:"runs"`
					} `json:"byline"`
					PauseText struct {
						SimpleText string `json:"simpleText"`
					} `json:"pauseText"`
					Background struct {
						Thumbnails []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"thumbnails"`
					} `json:"background"`
					CountDownSecs int `json:"countDownSecs"`
					CancelButton  struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
							Command struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										APIURL   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								GetSurveyCommand struct {
									Endpoint struct {
										Watch struct {
											Hack bool `json:"hack"`
										} `json:"watch"`
									} `json:"endpoint"`
									Action string `json:"action"`
								} `json:"getSurveyCommand"`
							} `json:"command"`
						} `json:"buttonRenderer"`
					} `json:"cancelButton"`
					NextButton struct {
						ButtonRenderer struct {
							Style              string `json:"style"`
							Size               string `json:"size"`
							IsDisabled         bool   `json:"isDisabled"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								WatchEndpoint struct {
									VideoID                            string `json:"videoId"`
									WatchEndpointSupportedOnesieConfig struct {
										HTML5PlaybackOnesieConfig struct {
											CommonConfig struct {
												URL string `json:"url"`
											} `json:"commonConfig"`
										} `json:"html5PlaybackOnesieConfig"`
									} `json:"watchEndpointSupportedOnesieConfig"`
								} `json:"watchEndpoint"`
							} `json:"navigationEndpoint"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
						} `json:"buttonRenderer"`
					} `json:"nextButton"`
					TrackingParams string `json:"trackingParams"`
					CloseButton    struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Icon       struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"closeButton"`
					ThumbnailOverlays []struct {
						ThumbnailOverlayTimeStatusRenderer struct {
							Text struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
							} `json:"text"`
							Style string `json:"style"`
							Icon  struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
						} `json:"thumbnailOverlayTimeStatusRenderer"`
					} `json:"thumbnailOverlays"`
					PreferImmediateRedirect      bool   `json:"preferImmediateRedirect"`
					VideoID                      string `json:"videoId"`
					WebShowNewAutonavCountdown   bool   `json:"webShowNewAutonavCountdown"`
					WebShowBigThumbnailEndscreen bool   `json:"webShowBigThumbnailEndscreen"`
					ShortViewCountText           struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
						Accessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
					} `json:"shortViewCountText"`
					CountDownSecsForFullscreen int `json:"countDownSecsForFullscreen"`
				} `json:"playerOverlayAutoplayRenderer"`
			} `json:"autoplay"`
			ShareButton struct {
				ButtonRenderer struct {
					Style      string `json:"style"`
					Size       string `json:"size"`
					IsDisabled bool   `json:"isDisabled"`
					Icon       struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								APIURL   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						ShareEntityServiceEndpoint struct {
							SerializedShareEntity string `json:"serializedShareEntity"`
							Commands              []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										UnifiedSharePanelRenderer struct {
											TrackingParams     string `json:"trackingParams"`
											ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
										} `json:"unifiedSharePanelRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
									BeReused  bool   `json:"beReused"`
								} `json:"openPopupAction"`
							} `json:"commands"`
						} `json:"shareEntityServiceEndpoint"`
					} `json:"navigationEndpoint"`
					Tooltip        string `json:"tooltip"`
					TrackingParams string `json:"trackingParams"`
				} `json:"buttonRenderer"`
			} `json:"shareButton"`
			AddToMenu struct {
				MenuRenderer struct {
					TrackingParams string `json:"trackingParams"`
				} `json:"menuRenderer"`
			} `json:"addToMenu"`
			VideoDetails struct {
				PlayerOverlayVideoDetailsRenderer struct {
					Title struct {
						SimpleText string `json:"simpleText"`
					} `json:"title"`
					Subtitle struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"subtitle"`
				} `json:"playerOverlayVideoDetailsRenderer"`
			} `json:"videoDetails"`
			LiveIndicatorText struct {
				SimpleText string `json:"simpleText"`
			} `json:"liveIndicatorText"`
			AutonavToggle struct {
				AutoplaySwitchButtonRenderer struct {
					OnEnabledCommand struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								APIURL   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SetSettingEndpoint struct {
							SettingItemID          string `json:"settingItemId"`
							BoolValue              bool   `json:"boolValue"`
							SettingItemIDForClient string `json:"settingItemIdForClient"`
						} `json:"setSettingEndpoint"`
					} `json:"onEnabledCommand"`
					OnDisabledCommand struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								APIURL   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SetSettingEndpoint struct {
							SettingItemID          string `json:"settingItemId"`
							BoolValue              bool   `json:"boolValue"`
							SettingItemIDForClient string `json:"settingItemIdForClient"`
						} `json:"setSettingEndpoint"`
					} `json:"onDisabledCommand"`
					EnabledAccessibilityData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"enabledAccessibilityData"`
					DisabledAccessibilityData struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"disabledAccessibilityData"`
					TrackingParams string `json:"trackingParams"`
					Enabled        bool   `json:"enabled"`
				} `json:"autoplaySwitchButtonRenderer"`
			} `json:"autonavToggle"`
			DecoratedPlayerBarRenderer struct {
				DecoratedPlayerBarRenderer struct {
				} `json:"decoratedPlayerBarRenderer"`
			} `json:"decoratedPlayerBarRenderer"`
			SpeedmasterUserEdu struct {
				SpeedmasterEduViewModel struct {
					BodyText struct {
						Content string `json:"content"`
					} `json:"bodyText"`
				} `json:"speedmasterEduViewModel"`
			} `json:"speedmasterUserEdu"`
			ShowPlaybackRateUpsellPanelCommand struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					InteractionLoggingCommandMetadata struct {
						ScreenVisualElement struct {
							UIType int `json:"uiType"`
						} `json:"screenVisualElement"`
					} `json:"interactionLoggingCommandMetadata"`
				} `json:"commandMetadata"`
				ShowDialogCommand struct {
					PanelLoadingStrategy struct {
						RequestTemplate struct {
							PanelID string `json:"panelId"`
							Params  string `json:"params"`
						} `json:"requestTemplate"`
						ScreenVe int `json:"screenVe"`
					} `json:"panelLoadingStrategy"`
				} `json:"showDialogCommand"`
			} `json:"showPlaybackRateUpsellPanelCommand"`
		} `json:"playerOverlayRenderer"`
	} `json:"playerOverlays"`
	OnResponseReceivedEndpoints []struct {
		ClickTrackingParams string `json:"clickTrackingParams"`
		CommandMetadata     struct {
			WebCommandMetadata struct {
				SendPost bool `json:"sendPost"`
			} `json:"webCommandMetadata"`
		} `json:"commandMetadata,omitempty"`
		SignalServiceEndpoint struct {
			Signal  string `json:"signal"`
			Actions []struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				SignalAction        struct {
					Signal string `json:"signal"`
				} `json:"signalAction"`
			} `json:"actions"`
		} `json:"signalServiceEndpoint,omitempty"`
		LoadMarkersCommand struct {
		} `json:"loadMarkersCommand,omitempty"`
	} `json:"onResponseReceivedEndpoints"`
	EngagementPanels []struct {
		EngagementPanelSectionListRenderer struct {
			Content struct {
				AdsEngagementPanelContentRenderer struct {
					Hack bool `json:"hack"`
				} `json:"adsEngagementPanelContentRenderer"`
			} `json:"content"`
			TargetID          string `json:"targetId"`
			Visibility        string `json:"visibility"`
			LoggingDirectives struct {
				TrackingParams string `json:"trackingParams"`
				Visibility     struct {
					Types string `json:"types"`
				} `json:"visibility"`
			} `json:"loggingDirectives"`
		} `json:"engagementPanelSectionListRenderer"`
	} `json:"engagementPanels"`
	Topbar struct {
		DesktopTopbarRenderer struct {
			Logo struct {
				TopbarLogoRenderer struct {
					IconImage struct {
						IconType string `json:"iconType"`
					} `json:"iconImage"`
					TooltipText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"tooltipText"`
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								APIURL      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseID string `json:"browseId"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					TrackingParams    string `json:"trackingParams"`
					OverrideEntityKey string `json:"overrideEntityKey"`
				} `json:"topbarLogoRenderer"`
			} `json:"logo"`
			Searchbox struct {
				FusionSearchboxRenderer struct {
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					PlaceholderText struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"placeholderText"`
					Config struct {
						WebSearchboxConfig struct {
							RequestLanguage     string `json:"requestLanguage"`
							RequestDomain       string `json:"requestDomain"`
							HasOnscreenKeyboard bool   `json:"hasOnscreenKeyboard"`
							FocusSearchbox      bool   `json:"focusSearchbox"`
						} `json:"webSearchboxConfig"`
					} `json:"config"`
					TrackingParams string `json:"trackingParams"`
					SearchEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SearchEndpoint struct {
							Query string `json:"query"`
						} `json:"searchEndpoint"`
					} `json:"searchEndpoint"`
					ClearButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Icon       struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
						} `json:"buttonRenderer"`
					} `json:"clearButton"`
				} `json:"fusionSearchboxRenderer"`
			} `json:"searchbox"`
			TrackingParams string `json:"trackingParams"`
			Interstitial   struct {
				ConsentBumpV2Renderer struct {
					InterstitialLogoAside struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"interstitialLogoAside"`
					LanguagePickerButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Icon struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"languagePickerButton"`
					InterstitialTitle struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"interstitialTitle"`
					CustomizeButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							TrackingParams string `json:"trackingParams"`
							Command        struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								URLEndpoint struct {
									URL string `json:"url"`
								} `json:"urlEndpoint"`
							} `json:"command"`
						} `json:"buttonRenderer"`
					} `json:"customizeButton"`
					AgreeButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
							TrackingParams string `json:"trackingParams"`
							Command        struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SaveConsentAction   struct {
									SocsCookie        string `json:"socsCookie"`
									SavePreferenceURL string `json:"savePreferenceUrl"`
								} `json:"saveConsentAction"`
							} `json:"command"`
						} `json:"buttonRenderer"`
					} `json:"agreeButton"`
					PrivacyLink struct {
						Runs []struct {
							Text               string `json:"text"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								URLEndpoint struct {
									URL string `json:"url"`
								} `json:"urlEndpoint"`
							} `json:"navigationEndpoint"`
						} `json:"runs"`
					} `json:"privacyLink"`
					TermsLink struct {
						Runs []struct {
							Text               string `json:"text"`
							NavigationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								URLEndpoint struct {
									URL string `json:"url"`
								} `json:"urlEndpoint"`
							} `json:"navigationEndpoint"`
						} `json:"runs"`
					} `json:"termsLink"`
					TrackingParams string `json:"trackingParams"`
					SignInButton   struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Icon struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							Tooltip        string `json:"tooltip"`
							TrackingParams string `json:"trackingParams"`
							Command        struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								SignInEndpoint struct {
									Hack bool `json:"hack"`
								} `json:"signInEndpoint"`
							} `json:"command"`
						} `json:"buttonRenderer"`
					} `json:"signInButton"`
					LanguageList struct {
						DropdownRenderer struct {
							Entries []struct {
								DropdownItemRenderer struct {
									Label struct {
										SimpleText string `json:"simpleText"`
									} `json:"label"`
									IsSelected      bool   `json:"isSelected"`
									StringValue     string `json:"stringValue"`
									OnSelectCommand struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												SendPost bool `json:"sendPost"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										SignalServiceEndpoint struct {
											Signal  string `json:"signal"`
											Actions []struct {
												ClickTrackingParams   string `json:"clickTrackingParams"`
												SelectLanguageCommand struct {
													Hl string `json:"hl"`
												} `json:"selectLanguageCommand"`
											} `json:"actions"`
										} `json:"signalServiceEndpoint"`
									} `json:"onSelectCommand"`
								} `json:"dropdownItemRenderer"`
							} `json:"entries"`
							Accessibility struct {
								Label string `json:"label"`
							} `json:"accessibility"`
						} `json:"dropdownRenderer"`
					} `json:"languageList"`
					ReadMoreButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							Icon struct {
								IconType string `json:"iconType"`
							} `json:"icon"`
							TrackingParams string `json:"trackingParams"`
							IconPosition   string `json:"iconPosition"`
						} `json:"buttonRenderer"`
					} `json:"readMoreButton"`
					DisableP13NButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								SimpleText string `json:"simpleText"`
							} `json:"text"`
							TrackingParams    string `json:"trackingParams"`
							AccessibilityData struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibilityData"`
							Command struct {
								ClickTrackingParams          string `json:"clickTrackingParams"`
								DisablePersonalizationAction struct {
									SocsCookie        string `json:"socsCookie"`
									SavePreferenceURL string `json:"savePreferenceUrl"`
								} `json:"disablePersonalizationAction"`
							} `json:"command"`
						} `json:"buttonRenderer"`
					} `json:"disableP13nButton"`
					LoadingMessage struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"loadingMessage"`
					ErrorMessage struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"errorMessage"`
					EomV1Text struct {
						EssentialCookieMsg struct {
							Begin struct {
								Runs []struct {
									Text               string `json:"text"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										URLEndpoint struct {
											URL string `json:"url"`
										} `json:"urlEndpoint"`
									} `json:"navigationEndpoint,omitempty"`
								} `json:"runs"`
							} `json:"begin"`
							Items []struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"items"`
						} `json:"essentialCookieMsg"`
						NonEssentialCookieMsg struct {
							Begin struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"begin"`
							Items []struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"items"`
						} `json:"nonEssentialCookieMsg"`
						IfReject struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"ifReject"`
						Personalization struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"personalization"`
						MoreOptions struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"moreOptions"`
					} `json:"eomV1Text"`
				} `json:"consentBumpV2Renderer"`
			} `json:"interstitial"`
			CountryCode   string `json:"countryCode"`
			TopbarButtons []struct {
				TopbarMenuButtonRenderer struct {
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					MenuRequest struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								APIURL   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								OpenPopupAction     struct {
									Popup struct {
										MultiPageMenuRenderer struct {
											TrackingParams     string `json:"trackingParams"`
											Style              string `json:"style"`
											ShowLoadingSpinner bool   `json:"showLoadingSpinner"`
										} `json:"multiPageMenuRenderer"`
									} `json:"popup"`
									PopupType string `json:"popupType"`
									BeReused  bool   `json:"beReused"`
								} `json:"openPopupAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"menuRequest"`
					TrackingParams string `json:"trackingParams"`
					Accessibility  struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
					Tooltip string `json:"tooltip"`
					Style   string `json:"style"`
				} `json:"topbarMenuButtonRenderer,omitempty"`
				ButtonRenderer struct {
					Style string `json:"style"`
					Size  string `json:"size"`
					Text  struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					Icon struct {
						IconType string `json:"iconType"`
					} `json:"icon"`
					NavigationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignInEndpoint struct {
							IdamTag string `json:"idamTag"`
						} `json:"signInEndpoint"`
					} `json:"navigationEndpoint"`
					TrackingParams string `json:"trackingParams"`
					TargetID       string `json:"targetId"`
				} `json:"buttonRenderer,omitempty"`
			} `json:"topbarButtons"`
			HotkeyDialog struct {
				HotkeyDialogRenderer struct {
					Title struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"title"`
					Sections []struct {
						HotkeyDialogSectionRenderer struct {
							Title struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"title"`
							Options []struct {
								HotkeyDialogSectionOptionRenderer struct {
									Label struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"label"`
									Hotkey string `json:"hotkey"`
								} `json:"hotkeyDialogSectionOptionRenderer"`
							} `json:"options"`
						} `json:"hotkeyDialogSectionRenderer"`
					} `json:"sections"`
					DismissButton struct {
						ButtonRenderer struct {
							Style      string `json:"style"`
							Size       string `json:"size"`
							IsDisabled bool   `json:"isDisabled"`
							Text       struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"text"`
							TrackingParams string `json:"trackingParams"`
						} `json:"buttonRenderer"`
					} `json:"dismissButton"`
					TrackingParams string `json:"trackingParams"`
				} `json:"hotkeyDialogRenderer"`
			} `json:"hotkeyDialog"`
			BackButton struct {
				ButtonRenderer struct {
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"backButton"`
			ForwardButton struct {
				ButtonRenderer struct {
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"forwardButton"`
			A11YSkipNavigationButton struct {
				ButtonRenderer struct {
					Style      string `json:"style"`
					Size       string `json:"size"`
					IsDisabled bool   `json:"isDisabled"`
					Text       struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool `json:"sendPost"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						SignalServiceEndpoint struct {
							Signal  string `json:"signal"`
							Actions []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								SignalAction        struct {
									Signal string `json:"signal"`
								} `json:"signalAction"`
							} `json:"actions"`
						} `json:"signalServiceEndpoint"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"a11ySkipNavigationButton"`
		} `json:"desktopTopbarRenderer"`
	} `json:"topbar"`
	PageVisualEffects []struct {
		CinematicContainerRenderer struct {
			GradientColorConfig []struct {
				DarkThemeColor int `json:"darkThemeColor"`
				StartLocation  int `json:"startLocation,omitempty"`
			} `json:"gradientColorConfig"`
			PresentationStyle string `json:"presentationStyle"`
			Config            struct {
				LightThemeBackgroundColor int64 `json:"lightThemeBackgroundColor"`
				DarkThemeBackgroundColor  int64 `json:"darkThemeBackgroundColor"`
				AnimationConfig           struct {
					CrossfadeStartOffset int `json:"crossfadeStartOffset"`
					MaxFrameRate         int `json:"maxFrameRate"`
				} `json:"animationConfig"`
				ColorSourceSizeMultiplier         float64 `json:"colorSourceSizeMultiplier"`
				ApplyClientImageBlur              bool    `json:"applyClientImageBlur"`
				BottomColorSourceHeightMultiplier float64 `json:"bottomColorSourceHeightMultiplier"`
				MaxBottomColorSourceHeight        int     `json:"maxBottomColorSourceHeight"`
				ColorSourceWidthMultiplier        float64 `json:"colorSourceWidthMultiplier"`
				ColorSourceHeightMultiplier       int     `json:"colorSourceHeightMultiplier"`
				BlurStrength                      int     `json:"blurStrength"`
				WatchFullscreenConfig             struct {
					ColorSourceWidthMultiplier  int     `json:"colorSourceWidthMultiplier"`
					ColorSourceHeightMultiplier int     `json:"colorSourceHeightMultiplier"`
					ScrimWidthMultiplier        float64 `json:"scrimWidthMultiplier"`
					ScrimHeightMultiplier       float64 `json:"scrimHeightMultiplier"`
					ScrimGradientConfig         struct {
						GradientType   string `json:"gradientType"`
						GradientColors []struct {
							LightThemeColor int `json:"lightThemeColor"`
							DarkThemeColor  int `json:"darkThemeColor"`
							StartLocation   int `json:"startLocation"`
						} `json:"gradientColors"`
						GradientStartPointX float64 `json:"gradientStartPointX"`
						GradientStartPointY float64 `json:"gradientStartPointY"`
						GradientEndPointX   int     `json:"gradientEndPointX"`
						GradientEndPointY   int     `json:"gradientEndPointY"`
					} `json:"scrimGradientConfig"`
				} `json:"watchFullscreenConfig"`
				EnableInLightTheme bool `json:"enableInLightTheme"`
			} `json:"config"`
		} `json:"cinematicContainerRenderer"`
	} `json:"pageVisualEffects"`
	FrameworkUpdates struct {
		EntityBatchUpdate struct {
			Mutations []struct {
				EntityKey string `json:"entityKey"`
				Type      string `json:"type"`
				Options   struct {
					PersistenceOption string `json:"persistenceOption"`
				} `json:"options,omitempty"`
				Payload struct {
					LikeStatusEntity struct {
						Key        string `json:"key"`
						LikeStatus string `json:"likeStatus"`
					} `json:"likeStatusEntity"`
				} `json:"payload,omitempty"`
			} `json:"mutations"`
			Timestamp struct {
				Seconds string `json:"seconds"`
				Nanos   int    `json:"nanos"`
			} `json:"timestamp"`
		} `json:"entityBatchUpdate"`
	} `json:"frameworkUpdates"`
}
