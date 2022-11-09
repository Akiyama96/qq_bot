package iservice

import (
	"QQ_bot/config"
	"QQ_bot/error_nt"
	"QQ_bot/internal/client"
	"QQ_bot/internal/dao"
	"QQ_bot/types"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	lulu = "21013446"

	myselfId = 1131568220
)

var lastTime int64

func SendNotification(info *types.LiveRoomInfo, times, beforeStatus int64, groupID int) {
	var msg string

	userInfo := client.GetCardInfo(info.Data.Uid)

	switch info.Data.LiveStatus {
	case 0:
		if beforeStatus == 1 {
			msg = fmt.Sprintf("%s下播啦!\n", userInfo.Data.Card.Name) +
				fmt.Sprintf("本次直播时间：%d小时%d分", times/3600, times%3600)
		}
	case 1:
		if beforeStatus == 0 || beforeStatus == 2 {
			msg = fmt.Sprintf("%s开播啦!\n", userInfo.Data.Card.Name) +
				fmt.Sprintf("直播间地址：https://live.bilibili.com/%d\n", info.Data.RoomId) +
				fmt.Sprintf("[CQ:image,file=%s]", info.Data.Keyframe)
		}
	case 2:
		if beforeStatus == 1 {
			msg = fmt.Sprintf("%s下播啦!\n", userInfo.Data.Card.Name) +
				fmt.Sprintf("本次直播时间：%d小时%d分\n", times/3600, times%3600)
		}
		msg += fmt.Sprintf("%s轮播中\n", userInfo.Data.Card.Name) +
			fmt.Sprintf("直播间地址：https://live.bilibili.com/%d\n", info.Data.RoomId) +
			fmt.Sprintf("[CQ:image,file=%s]", info.Data.Keyframe)
	}

	err := client.SendNotificationMsg("private", myselfId, msg)
	if err != nil {
		log.Println(err)
		error_nt.SendInfo(err.Error())
	}

	err = client.SendNotificationMsg("group", groupID, msg)
	if err != nil {
		log.Println(err)
		error_nt.SendInfo(err.Error())
	}
}

func HandleEvent(data *types.Event) {
	switch data.PostType {
	case "message":
		log.Println(fmt.Sprintf("Rev msg from %s:%s", data.Sender.Nickname, data.Message))
		handleMsg(data)
	case "request":
		if data.SubType == "add" || data.SubType == "invite" {
			err := client.AppendInvite(data)
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "meta_event":

	}
}

func handleMsg(data *types.Event) {

	var id int

	if data.MessageType == "group" {
		id = data.GroupId
	} else if data.MessageType == "private" {
		id = data.UserId
	}

	if len(data.Message) > 11 && data.Message[:3] == "cmd" && data.Sender.UserId == config.Srv.Bot.QQ {
		switch data.Message[4:9] {
		case "srvcg":
			cmd := []byte(data.Message[10:])
			serviceInfo := types.BilibiliService{}
			err := json.Unmarshal(cmd, &serviceInfo)
			if err != nil {
				error_nt.SendInfo(err.Error())
				log.Println(err)
			}

			creatNewService(serviceInfo)
		}
	}

	if len(data.Message) > 12 {
		switch data.Message[0:12] {
		case "来点色图":
			if (time.Now().Unix() - lastTime) < 10 {
				err := client.SendNotificationMsg(data.MessageType, id, fmt.Sprintf("CD中...%d", 10-(time.Now().Unix()-lastTime)))
				if err != nil {
					log.Println(err)
				}
			}
			go func() {
				lastTime = time.Now().Unix()
				SendHPic(data.Message[13:], id, data.MessageType)
			}()

			err := client.SendNotificationMsg(data.MessageType, id, "正在找色图")
			if err != nil {
				log.Println(err)
			}
		}
	}

	if data.MessageType == "group" {
		srvInfo := dao.GetServiceInfoByGroupID(data.GroupId)

		if srvInfo == nil {
			return
		}

		switch data.Message {
		case "--help":
			var msg string

			var functions = []string{"直播推送", "动态推送"}

			var cmds = []string{"fans", "last_dynamic", "live_status"}

			msg = "——bot 当前支持的命令——\n\n"

			for _, cmd := range cmds {
				msg += cmd + "\n"
			}

			msg += "\n——bot 当前支持的功能——\n\n"

			for _, function := range functions {
				msg += function + "\n"
			}

			err := client.SendNotificationMsg(data.MessageType, id, msg)
			if err != nil {
				log.Println(err)
			}

		case "fans":
			info := client.GetUserStatInfo(strconv.Itoa(srvInfo.UserID))

			userInfo := client.GetCardInfo(srvInfo.UserID)

			msg := fmt.Sprintf("%s 当前粉丝量：%d", userInfo.Data.Card.Name, info.Data.Follower)

			err := client.SendNotificationMsg(data.MessageType, id, msg)
			if err != nil {
				log.Println(err)
			}

		case "last_dynamic":
			var flag int

			info := client.GetSpaceInfo(strconv.Itoa(srvInfo.UserID))

			if info.Data.Items[0].Modules.ModuleTag.Text == "置顶" {
				flag = 1
			}

			SendSpaceMsg(info, data.MessageType, id, flag)

		case "live_status":
			var msg string
			userInfo := client.GetCardInfo(srvInfo.UserID)
			info := client.GetLiveRoomInfo(strconv.Itoa(srvInfo.RoomID))
			switch info.Data.LiveStatus {
			case 0:
				msg = fmt.Sprintf("%s未在直播!\n", userInfo.Data.Card.Name)
				//fmt.Sprintf("本次直播时间：%d小时%d分", times/3600, times%3600)
			case 1:
				msg = fmt.Sprintf("%s直播中!\n", userInfo.Data.Card.Name) +
					fmt.Sprintf("直播间地址：https://live.bilibili.com/%d\n", info.Data.RoomId) +
					fmt.Sprintf("[CQ:image,file=%s]", info.Data.Keyframe)
			case 2:
				msg = fmt.Sprintf("%s轮播中!\n", userInfo.Data.Card.Name) +
					fmt.Sprintf("直播间地址：https://live.bilibili.com/%d\n", info.Data.RoomId) +
					fmt.Sprintf("[CQ:image,file=%s]", info.Data.Keyframe)
			}

			err := client.SendNotificationMsg(data.MessageType, id, msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func SendSpaceMsg(info *types.SpaceInfo, target string, id, flag int) {
	var msg, topic, url, text, pic string
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	topic = info.Data.Items[flag].Modules.ModuleAuthor.Name + " " + info.Data.Items[flag].Modules.ModuleAuthor.PubAction
	url = info.Data.Items[flag].Modules.ModuleAuthor.JumpUrl
	if info.Data.Items[flag].Modules.ModuleDynamic.Desc != nil {
		text = info.Data.Items[flag].Modules.ModuleDynamic.Desc.Text
	}

	if info.Data.Items[flag].Modules.ModuleDynamic.Major != nil {
		pic = info.Data.Items[flag].Modules.ModuleDynamic.Major.Article.Covers[0]
		pic = fmt.Sprintf("[CQ:image,file=%s]", pic)
	}

	msg = fmt.Sprintf(topic + "\n" + url + "\n" + text + "\n" + pic)

	err := client.SendNotificationMsg(target, id, msg)
	if err != nil {
		log.Println(err)
		return
	}
}

func SendHPic(tag string, id int, target string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	var tags = make([]string, 0)

	for i, v := range tag {
		if v == ',' {
			if i != 0 {
				tags = append(tags, tag[0:i])
			}
		}
	}

	if len(tags) == 0 {
		tags = append(tags, tag)
	}
	picInfo := client.GetPic(tags)

	var msg string

	if len(picInfo.Data) > 0 {
		msg = fmt.Sprintf("标题：%s \n PID:%d\n ", picInfo.Data[0].Artwork.Title, picInfo.Data[0].Artwork.Id)
		msg += fmt.Sprintf("Medium:%s\n ", picInfo.Data[0].Urls.Medium)
		//msg += fmt.Sprintf("Large:%s\n ", picInfo.Data[0].Urls.Large)
		//msg += fmt.Sprintf("Medium:%s\n ", picInfo.Data[0].Urls.Medium)
		//msg += fmt.Sprintf("[CQ:image,file=%s]", picInfo.Data[0].Urls.Medium)
	} else {
		msg = "没有找到对应tag的图片"
	}

	err := client.SendNotificationMsg(target, id, msg)
	if err != nil {
		log.Println(err)
	}
}

func creatNewService(serviceInfo types.BilibiliService) {
	dao.UpdateBilibiliServiceInfo(serviceInfo)

	if serviceInfo.LiveNotification == 1 {
		NewLiveServiceAdd(serviceInfo)
	} else if serviceInfo.LiveNotification == 0 {
		CancelService(serviceInfo.RoomID, serviceInfo.GroupID)
	}

	if serviceInfo.SpaceNotification == 1 {
		NewSpaceServiceAdd(serviceInfo)
	} else if serviceInfo.SpaceNotification == 0 {
		CancelService(serviceInfo.UserID, serviceInfo.GroupID)
	}
}
