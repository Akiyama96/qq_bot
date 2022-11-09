package types

type LiveRoomInfo struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		Uid              int      `json:"uid"`
		RoomId           int      `json:"room_id"`
		ShortId          int      `json:"short_id"`
		Attention        int      `json:"attention"`
		Online           int      `json:"online"`
		IsPortrait       bool     `json:"is_portrait"`
		Description      string   `json:"description"`
		LiveStatus       int      `json:"live_status"`
		AreaId           int      `json:"area_id"`
		ParentAreaId     int      `json:"parent_area_id"`
		ParentAreaName   string   `json:"parent_area_name"`
		OldAreaId        int      `json:"old_area_id"`
		Background       string   `json:"background"`
		Title            string   `json:"title"`
		UserCover        string   `json:"user_cover"`
		Keyframe         string   `json:"keyframe"`
		IsStrictRoom     bool     `json:"is_strict_room"`
		LiveTime         string   `json:"live_time"`
		Tags             string   `json:"tags"`
		IsAnchor         int      `json:"is_anchor"`
		RoomSilentType   string   `json:"room_silent_type"`
		RoomSilentLevel  int      `json:"room_silent_level"`
		RoomSilentSecond int      `json:"room_silent_second"`
		AreaName         string   `json:"area_name"`
		Pendants         string   `json:"pendants"`
		AreaPendants     string   `json:"area_pendants"`
		HotWords         []string `json:"hot_words"`
		HotWordsStatus   int      `json:"hot_words_status"`
		Verify           string   `json:"verify"`
		NewPendants      struct {
			Frame struct {
				Name       string `json:"name"`
				Value      string `json:"value"`
				Position   int    `json:"position"`
				Desc       string `json:"desc"`
				Area       int    `json:"area"`
				AreaOld    int    `json:"area_old"`
				BgColor    string `json:"bg_color"`
				BgPic      string `json:"bg_pic"`
				UseOldArea bool   `json:"use_old_area"`
			} `json:"frame"`
			Badge struct {
				Name     string `json:"name"`
				Position int    `json:"position"`
				Value    string `json:"value"`
				Desc     string `json:"desc"`
			} `json:"badge"`
			MobileFrame struct {
				Name       string `json:"name"`
				Value      string `json:"value"`
				Position   int    `json:"position"`
				Desc       string `json:"desc"`
				Area       int    `json:"area"`
				AreaOld    int    `json:"area_old"`
				BgColor    string `json:"bg_color"`
				BgPic      string `json:"bg_pic"`
				UseOldArea bool   `json:"use_old_area"`
			} `json:"mobile_frame"`
			MobileBadge interface{} `json:"mobile_badge"`
		} `json:"new_pendants"`
		UpSession            string `json:"up_session"`
		PkStatus             int    `json:"pk_status"`
		PkId                 int    `json:"pk_id"`
		BattleId             int    `json:"battle_id"`
		AllowChangeAreaTime  int    `json:"allow_change_area_time"`
		AllowUploadCoverTime int    `json:"allow_upload_cover_time"`
		StudioInfo           struct {
			Status     int           `json:"status"`
			MasterList []interface{} `json:"master_list"`
		} `json:"studio_info"`
	} `json:"data"`
}

type UserStatInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Mid       int `json:"mid"`
		Following int `json:"following"`
		Whisper   int `json:"whisper"`
		Black     int `json:"black"`
		Follower  int `json:"follower"`
	} `json:"data"`
}

type SpaceInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		HasMore bool `json:"has_more"`
		Items   []struct {
			Basic struct {
				CommentIdStr string `json:"comment_id_str"`
				CommentType  int    `json:"comment_type"`
				LikeIcon     struct {
					ActionUrl string `json:"action_url"`
					EndUrl    string `json:"end_url"`
					Id        int    `json:"id"`
					StartUrl  string `json:"start_url"`
				} `json:"like_icon"`
				RidStr string `json:"rid_str"`
			} `json:"basic"`
			IdStr   string `json:"id_str"`
			Modules struct {
				ModuleAuthor struct {
					Decorate struct {
						CardUrl string `json:"card_url"`
						Fan     struct {
							Color  string `json:"color"`
							IsFan  bool   `json:"is_fan"`
							NumStr string `json:"num_str"`
							Number int    `json:"number"`
						} `json:"fan"`
						Id      int    `json:"id"`
						JumpUrl string `json:"jump_url"`
						Name    string `json:"name"`
						Type    int    `json:"type"`
					} `json:"decorate"`
					Face           string      `json:"face"`
					FaceNft        bool        `json:"face_nft"`
					Following      interface{} `json:"following"`
					JumpUrl        string      `json:"jump_url"`
					Label          string      `json:"label"`
					Mid            int         `json:"mid"`
					Name           string      `json:"name"`
					OfficialVerify struct {
						Desc string `json:"desc"`
						Type int    `json:"type"`
					} `json:"official_verify"`
					Pendant struct {
						Expire            int    `json:"expire"`
						Image             string `json:"image"`
						ImageEnhance      string `json:"image_enhance"`
						ImageEnhanceFrame string `json:"image_enhance_frame"`
						Name              string `json:"name"`
						Pid               int    `json:"pid"`
					} `json:"pendant"`
					PubAction       string `json:"pub_action"`
					PubLocationText string `json:"pub_location_text"`
					PubTime         string `json:"pub_time"`
					PubTs           int    `json:"pub_ts"`
					Type            string `json:"type"`
					Vip             struct {
						AvatarSubscript    int    `json:"avatar_subscript"`
						AvatarSubscriptUrl string `json:"avatar_subscript_url"`
						DueDate            int64  `json:"due_date"`
						Label              struct {
							BgColor               string `json:"bg_color"`
							BgStyle               int    `json:"bg_style"`
							BorderColor           string `json:"border_color"`
							ImgLabelUriHans       string `json:"img_label_uri_hans"`
							ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
							ImgLabelUriHant       string `json:"img_label_uri_hant"`
							ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
							LabelTheme            string `json:"label_theme"`
							Path                  string `json:"path"`
							Text                  string `json:"text"`
							TextColor             string `json:"text_color"`
							UseImgLabel           bool   `json:"use_img_label"`
						} `json:"label"`
						NicknameColor string `json:"nickname_color"`
						Status        int    `json:"status"`
						ThemeType     int    `json:"theme_type"`
						Type          int    `json:"type"`
					} `json:"vip"`
				} `json:"module_author"`
				ModuleDynamic struct {
					Additional interface{} `json:"additional"`
					Desc       *struct {
						RichTextNodes []struct {
							OrigText string `json:"orig_text"`
							Text     string `json:"text"`
							Type     string `json:"type"`
							Emoji    struct {
								IconUrl string `json:"icon_url"`
								Size    int    `json:"size"`
								Text    string `json:"text"`
								Type    int    `json:"type"`
							} `json:"emoji,omitempty"`
						} `json:"rich_text_nodes"`
						Text string `json:"text"`
					} `json:"desc"`
					Major *struct {
						Article struct {
							Covers  []string `json:"covers"`
							Desc    string   `json:"desc"`
							Id      int      `json:"id"`
							JumpUrl string   `json:"jump_url"`
							Label   string   `json:"label"`
							Title   string   `json:"title"`
						} `json:"article,omitempty"`
						Type string `json:"type"`
						Draw struct {
							Id    int `json:"id"`
							Items []struct {
								Height int           `json:"height"`
								Size   float64       `json:"size"`
								Src    string        `json:"src"`
								Tags   []interface{} `json:"tags"`
								Width  int           `json:"width"`
							} `json:"items"`
						} `json:"draw,omitempty"`
					} `json:"major"`
					Topic interface{} `json:"topic"`
				} `json:"module_dynamic"`
				ModuleInteraction struct {
					Items []struct {
						Desc struct {
							RichTextNodes []struct {
								OrigText string `json:"orig_text"`
								Rid      string `json:"rid,omitempty"`
								Text     string `json:"text"`
								Type     string `json:"type"`
								Emoji    struct {
									IconUrl string `json:"icon_url"`
									Size    int    `json:"size"`
									Text    string `json:"text"`
									Type    int    `json:"type"`
								} `json:"emoji,omitempty"`
							} `json:"rich_text_nodes"`
							Text string `json:"text"`
						} `json:"desc"`
						Type int `json:"type"`
					} `json:"items"`
				} `json:"module_interaction"`
				ModuleMore struct {
					ThreePointItems []struct {
						Label string `json:"label"`
						Type  string `json:"type"`
					} `json:"three_point_items"`
				} `json:"module_more"`
				ModuleStat struct {
					Comment struct {
						Count     int  `json:"count"`
						Forbidden bool `json:"forbidden"`
					} `json:"comment"`
					Forward struct {
						Count     int  `json:"count"`
						Forbidden bool `json:"forbidden"`
					} `json:"forward"`
					Like struct {
						Count     int  `json:"count"`
						Forbidden bool `json:"forbidden"`
						Status    bool `json:"status"`
					} `json:"like"`
				} `json:"module_stat"`
				ModuleTag struct {
					Text string `json:"text"`
				} `json:"module_tag,omitempty"`
			} `json:"modules"`
			Type    string `json:"type"`
			Visible bool   `json:"visible"`
			Orig    struct {
				Basic struct {
					CommentIdStr string `json:"comment_id_str"`
					CommentType  int    `json:"comment_type"`
					LikeIcon     struct {
						ActionUrl string `json:"action_url"`
						EndUrl    string `json:"end_url"`
						Id        int    `json:"id"`
						StartUrl  string `json:"start_url"`
					} `json:"like_icon"`
					RidStr string `json:"rid_str"`
				} `json:"basic"`
				IdStr   string `json:"id_str"`
				Modules struct {
					ModuleAuthor struct {
						Decorate struct {
							CardUrl string `json:"card_url"`
							Fan     struct {
								Color  string `json:"color"`
								IsFan  bool   `json:"is_fan"`
								NumStr string `json:"num_str"`
								Number int    `json:"number"`
							} `json:"fan"`
							Id      int    `json:"id"`
							JumpUrl string `json:"jump_url"`
							Name    string `json:"name"`
							Type    int    `json:"type"`
						} `json:"decorate"`
						Face           string      `json:"face"`
						FaceNft        bool        `json:"face_nft"`
						Following      interface{} `json:"following"`
						JumpUrl        string      `json:"jump_url"`
						Label          string      `json:"label"`
						Mid            int         `json:"mid"`
						Name           string      `json:"name"`
						OfficialVerify struct {
							Desc string `json:"desc"`
							Type int    `json:"type"`
						} `json:"official_verify"`
						Pendant struct {
							Expire            int    `json:"expire"`
							Image             string `json:"image"`
							ImageEnhance      string `json:"image_enhance"`
							ImageEnhanceFrame string `json:"image_enhance_frame"`
							Name              string `json:"name"`
							Pid               int    `json:"pid"`
						} `json:"pendant"`
						PubAction string `json:"pub_action"`
						PubTime   string `json:"pub_time"`
						PubTs     int    `json:"pub_ts"`
						Type      string `json:"type"`
						Vip       struct {
							AvatarSubscript    int    `json:"avatar_subscript"`
							AvatarSubscriptUrl string `json:"avatar_subscript_url"`
							DueDate            int64  `json:"due_date"`
							Label              struct {
								BgColor               string `json:"bg_color"`
								BgStyle               int    `json:"bg_style"`
								BorderColor           string `json:"border_color"`
								ImgLabelUriHans       string `json:"img_label_uri_hans"`
								ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
								ImgLabelUriHant       string `json:"img_label_uri_hant"`
								ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
								LabelTheme            string `json:"label_theme"`
								Path                  string `json:"path"`
								Text                  string `json:"text"`
								TextColor             string `json:"text_color"`
								UseImgLabel           bool   `json:"use_img_label"`
							} `json:"label"`
							NicknameColor string `json:"nickname_color"`
							Status        int    `json:"status"`
							ThemeType     int    `json:"theme_type"`
							Type          int    `json:"type"`
						} `json:"vip"`
					} `json:"module_author"`
					ModuleDynamic struct {
						Additional interface{} `json:"additional"`
						Desc       struct {
							RichTextNodes []struct {
								OrigText string `json:"orig_text"`
								Rid      string `json:"rid,omitempty"`
								Text     string `json:"text"`
								Type     string `json:"type"`
								Emoji    struct {
									IconUrl string `json:"icon_url"`
									Size    int    `json:"size"`
									Text    string `json:"text"`
									Type    int    `json:"type"`
								} `json:"emoji,omitempty"`
								JumpUrl string `json:"jump_url,omitempty"`
							} `json:"rich_text_nodes"`
							Text string `json:"text"`
						} `json:"desc"`
						Major struct {
							Draw struct {
								Id    int `json:"id"`
								Items []struct {
									Height int           `json:"height"`
									Size   float64       `json:"size"`
									Src    string        `json:"src"`
									Tags   []interface{} `json:"tags"`
									Width  int           `json:"width"`
								} `json:"items"`
							} `json:"draw,omitempty"`
							Type    string `json:"type"`
							Archive struct {
								Aid   string `json:"aid"`
								Badge struct {
									BgColor string `json:"bg_color"`
									Color   string `json:"color"`
									Text    string `json:"text"`
								} `json:"badge"`
								Bvid           string `json:"bvid"`
								Cover          string `json:"cover"`
								Desc           string `json:"desc"`
								DisablePreview int    `json:"disable_preview"`
								DurationText   string `json:"duration_text"`
								JumpUrl        string `json:"jump_url"`
								Stat           struct {
									Danmaku string `json:"danmaku"`
									Play    string `json:"play"`
								} `json:"stat"`
								Title string `json:"title"`
								Type  int    `json:"type"`
							} `json:"archive,omitempty"`
						} `json:"major"`
						Topic interface{} `json:"topic"`
					} `json:"module_dynamic"`
				} `json:"modules"`
				Type    string `json:"type"`
				Visible bool   `json:"visible"`
			} `json:"orig,omitempty"`
		} `json:"items"`
		Offset         string `json:"offset"`
		UpdateBaseline string `json:"update_baseline"`
		UpdateNum      int    `json:"update_num"`
	} `json:"data"`
}

type BilibiliService struct {
	Name              string `json:"name"`
	UserID            int    `json:"user_id"`
	GroupID           int    `json:"group_id"`
	RoomID            int    `json:"room_id"`
	LiveNotification  int    `json:"live_notification"`
	SpaceNotification int    `json:"space_notification"`
}

type BilibiliUserInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Card struct {
			Mid         string        `json:"mid"`
			Name        string        `json:"name"`
			Approve     bool          `json:"approve"`
			Sex         string        `json:"sex"`
			Rank        string        `json:"rank"`
			Face        string        `json:"face"`
			FaceNft     int           `json:"face_nft"`
			FaceNftType int           `json:"face_nft_type"`
			DisplayRank string        `json:"DisplayRank"`
			Regtime     int           `json:"regtime"`
			Spacesta    int           `json:"spacesta"`
			Birthday    string        `json:"birthday"`
			Place       string        `json:"place"`
			Description string        `json:"description"`
			Article     int           `json:"article"`
			Attentions  []interface{} `json:"attentions"`
			Fans        int           `json:"fans"`
			Friend      int           `json:"friend"`
			Attention   int           `json:"attention"`
			Sign        string        `json:"sign"`
			LevelInfo   struct {
				CurrentLevel int `json:"current_level"`
				CurrentMin   int `json:"current_min"`
				CurrentExp   int `json:"current_exp"`
				NextExp      int `json:"next_exp"`
			} `json:"level_info"`
			Pendant struct {
				Pid               int    `json:"pid"`
				Name              string `json:"name"`
				Image             string `json:"image"`
				Expire            int    `json:"expire"`
				ImageEnhance      string `json:"image_enhance"`
				ImageEnhanceFrame string `json:"image_enhance_frame"`
			} `json:"pendant"`
			Nameplate struct {
				Nid        int    `json:"nid"`
				Name       string `json:"name"`
				Image      string `json:"image"`
				ImageSmall string `json:"image_small"`
				Level      string `json:"level"`
				Condition  string `json:"condition"`
			} `json:"nameplate"`
			Official struct {
				Role  int    `json:"role"`
				Title string `json:"title"`
				Desc  string `json:"desc"`
				Type  int    `json:"type"`
			} `json:"Official"`
			OfficialVerify struct {
				Type int    `json:"type"`
				Desc string `json:"desc"`
			} `json:"official_verify"`
			Vip struct {
				Type       int   `json:"type"`
				Status     int   `json:"status"`
				DueDate    int64 `json:"due_date"`
				VipPayType int   `json:"vip_pay_type"`
				ThemeType  int   `json:"theme_type"`
				Label      struct {
					Path                  string `json:"path"`
					Text                  string `json:"text"`
					LabelTheme            string `json:"label_theme"`
					TextColor             string `json:"text_color"`
					BgStyle               int    `json:"bg_style"`
					BgColor               string `json:"bg_color"`
					BorderColor           string `json:"border_color"`
					UseImgLabel           bool   `json:"use_img_label"`
					ImgLabelUriHans       string `json:"img_label_uri_hans"`
					ImgLabelUriHant       string `json:"img_label_uri_hant"`
					ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
					ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`
				NicknameColor      string `json:"nickname_color"`
				Role               int    `json:"role"`
				AvatarSubscriptUrl string `json:"avatar_subscript_url"`
				TvVipStatus        int    `json:"tv_vip_status"`
				TvVipPayType       int    `json:"tv_vip_pay_type"`
				VipType            int    `json:"vipType"`
				VipStatus          int    `json:"vipStatus"`
			} `json:"vip"`
			IsSeniorMember int `json:"is_senior_member"`
		} `json:"card"`
		Following    bool `json:"following"`
		ArchiveCount int  `json:"archive_count"`
		ArticleCount int  `json:"article_count"`
		Follower     int  `json:"follower"`
		LikeNum      int  `json:"like_num"`
	} `json:"data"`
}
