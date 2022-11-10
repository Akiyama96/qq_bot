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
	"sync"
	"time"
)

var tarotTime sync.Map

var lastTime int64

func SendNotification(info *types.LiveRoomInfo, times, beforeStatus int64, groupID int) {
	var msg string

	userInfo := client.GetCardInfo(info.Data.Uid)

	switch info.Data.LiveStatus {
	case 0:
		if beforeStatus == 1 {
			msg = fmt.Sprintf("%s下播啦!\n", userInfo.Data.Card.Name) +
				fmt.Sprintf("本次直播时间：%d小时%d分", times/3600, (times%3600)/60)
		}
	case 1:
		if beforeStatus == 0 || beforeStatus == 2 {
			msg = fmt.Sprintf("%s开播啦!\n", userInfo.Data.Card.Name) +
				fmt.Sprintf("直播间地址：https://live.bilibili.com/%d\n", info.Data.RoomId) +
				fmt.Sprintf("当前有%d人正在观看~\n", info.Data.Online) +
				fmt.Sprintf("[CQ:image,file=%s]", info.Data.UserCover)
		}
	case 2:
		if beforeStatus == 1 {
			msg = fmt.Sprintf("%s下播啦!\n", userInfo.Data.Card.Name) +
				fmt.Sprintf("本次直播时间：%d小时%d分\n", times/3600, (times%3600)/60)
		}
		msg += fmt.Sprintf("%s轮播中\n", userInfo.Data.Card.Name) +
			fmt.Sprintf("直播间地址：https://live.bilibili.com/%d\n", info.Data.RoomId) +
			fmt.Sprintf("[CQ:image,file=%s]", info.Data.Keyframe)
	}

	err := client.SendNotificationMsg("private", config.Srv.Bot.QQ, msg)
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

	if len(data.Message) > 23 {
		if data.Message[0:21] == "[CQ:at,qq=1497312823]" {
			client.Xiaoai(data.MessageType, id, data.Message)
			//err := client.SendNotificationMsg(data.MessageType, id, "确实")
			//if err != nil {
			//	log.Println(err)
			//}
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
		case "&amp;抽卡":
			key := data.Sender.UserId

			if v, ok := tarotTime.Load(key); ok {
				if (time.Now().Unix() - v.(int64)) < (60 * 60 * 24) {
					err := client.SendNotificationMsg(data.MessageType, id, "你今天已经抽过卡啦~")
					if err != nil {
						log.Println(err)
					}
					return
				}
			}

			tarotTime.Store(key, time.Now().Unix())

			tarot(data)
		case "--help":
			var msg string

			var functions = []string{"直播推送", "动态推送"}

			var cmds = []string{"&粉丝数", "&查看动态", "&直播状态", "&抽卡", "&被隐藏的功能"}

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

		case "&amp;粉丝数":
			info := client.GetUserStatInfo(strconv.Itoa(srvInfo.UserID))

			userInfo := client.GetCardInfo(srvInfo.UserID)

			msg := fmt.Sprintf("%s 当前粉丝量：%d", userInfo.Data.Card.Name, info.Data.Follower)

			err := client.SendNotificationMsg(data.MessageType, id, msg)
			if err != nil {
				log.Println(err)
			}

		case "&amp;查看动态":
			var flag int

			info := client.GetSpaceInfo(strconv.Itoa(srvInfo.UserID))

			if info.Data.Items[0].Modules.ModuleTag.Text == "置顶" {
				flag = 1
			}

			SendSpaceMsg(info, data.MessageType, id, flag)

		case "&amp;直播状态":
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
					fmt.Sprintf("当前有%d人正在观看~\n", info.Data.Online) +
					fmt.Sprintf("[CQ:image,file=%s]", info.Data.UserCover)
			case 2:
				msg = fmt.Sprintf("%s轮播中!\n", userInfo.Data.Card.Name) +
					fmt.Sprintf("直播间地址：https://live.bilibili.com/%d\n", info.Data.RoomId) +
					fmt.Sprintf("[CQ:image,file=%s]", info.Data.UserCover)
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

	var r18 int

	if id == 1131569220 {
		r18 = 1
	}

	picInfo := client.GetPic(tags, r18)

	var msg string

	if picInfo.Detail == "色图库中没找到色图~" {
		msg = "色图库中没找到色图~"
		err := client.SendNotificationMsg(target, id, msg)
		if err != nil {
			log.Println(err)
		}
		return
	}

	if len(picInfo.Data) > 0 {
		msg = fmt.Sprintf("标题：%s \n PID:%d\n ", picInfo.Data[0].Artwork.Title, picInfo.Data[0].Artwork.Id)
		msg += fmt.Sprintf("Medium:%s\n ", picInfo.Data[0].Urls.Original)
		//msg += fmt.Sprintf("Large:%s\n ", picInfo.Data[0].Urls.Large)
		//msg += fmt.Sprintf("Medium:%s\n ", picInfo.Data[0].Urls.Medium)
		msg += fmt.Sprintf("[CQ:image,file=%s]", picInfo.Data[0].Urls.Medium)
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
