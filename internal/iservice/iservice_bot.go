package iservice

import (
	"QQ_bot/error_nt"
	"QQ_bot/internal/client"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

func NotificationService(ctx context.Context, roomID, groupID int) {
	roomIDString := strconv.Itoa(roomID)
	var liveStatus, times, flag int64

	for {
		select {
		case <-ctx.Done():
			error_nt.SendInfo(fmt.Sprintf(
				"Service \nRoomID:%d,\nGroupID:%d\n has been stopped.",
				roomID,
				groupID,
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
				SendNotification(info, times, liveStatus, groupID)
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
	var last string
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

		if len(info.Data.Items) > 1 {

			if info.Data.Items[0].Modules.ModuleTag.Text == "置顶" {
				flag1 = 1
			} else {
				flag1 = 0
			}

			if flag > 0 {
				if info.Data.Items[flag1].Modules.ModuleDynamic.Desc != nil && info.Data.Items[flag].Modules.ModuleDynamic.Desc.Text != last {
					SendSpaceMsg(info, "group", groupID, flag1)
				}
			}

			last = info.Data.Items[flag1].Modules.ModuleDynamic.Desc.Text
		}

		flag = 1
		time.Sleep(30 * time.Second)
	}
}
