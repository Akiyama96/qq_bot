package iservice

import (
	"QQ_bot/error_nt"
	"QQ_bot/internal/client"
	"QQ_bot/types"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

func NotificationService(ctx context.Context, service types.BilibiliService) {
	roomIDString := strconv.Itoa(service.RoomID)
	var liveStatus, times, flag int64

	for {
		select {
		case <-ctx.Done():
			error_nt.SendInfo(fmt.Sprintf(
				"Service \nRoomID:%d,\nGroupID:%d\n has been stopped.",
				service.RoomID,
				service.GroupID,
			))
			return
		default:
		}

		info := client.GetLiveRoomInfo(roomIDString)

		if info == nil || info.Msg != "ok" {
			time.Sleep(2 * time.Second)
			continue
		}

		if liveStatus != int64(info.Data.LiveStatus) {
			//计算直播时间
			if times != 0 && (info.Data.LiveStatus == 0 || info.Data.LiveStatus == 2) {
				times = time.Now().Unix() - times
			}

			if flag == 1 {
				SendNotification(info, times, liveStatus, service.GroupID, service.AtAll)
			}

			liveStatus = int64(info.Data.LiveStatus)
			times = time.Now().Unix()
		}

		flag = 1
		time.Sleep(2 * time.Second)
	}
}

func Space(ctx context.Context, userID, groupID int) {
	var flag, flag1 int
	var userIDString = strconv.Itoa(userID)
	var last []byte
	for {
		select {
		case <-ctx.Done():
			error_nt.SendInfo(fmt.Sprintf(
				"Space service \nUserID:%d,\nGroupID:%d\n has been stopped.",
				userID,
				groupID,
			))
			return
		default:
		}

		info := client.GetSpaceInfo(userIDString)

		if info == nil || info.Code != 0 {
			log.Println("get space info failed")
			time.Sleep(30 * time.Second)
			continue
		}

		if len(info.Data.Items) > 2 {
			if info.Data.Items[0].Modules.ModuleTag.Text == "置顶" {
				flag1 = 1
			} else {
				flag1 = 0
			}

			if flag > 0 && info.Data.Items[flag1].Modules.ModuleDynamic.Desc != nil {
				last1, _ := json.Marshal(info.Data.Items[flag1].Modules.ModuleDynamic.Desc)
				if len(last1) != len(last) {
					SendSpaceMsg(info, "group", groupID, flag1)
				}
			}

			if info.Data.Items[flag1].Modules.ModuleDynamic.Desc != nil {
				last, _ = json.Marshal(info.Data.Items[flag1].Modules.ModuleDynamic.Desc)
				flag = 1
			}
		}

		time.Sleep(30 * time.Second)
	}
}
