package iservice

import (
	"QQ_bot/config"
	"QQ_bot/error_nt"
	"QQ_bot/internal/client"
	"QQ_bot/types"
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
)

// 存储对应的上下文取消函数
var ctxMap sync.Map

func NewLiveServiceAdd(service types.BilibiliService) {
	var err error
	key := strconv.Itoa(service.RoomID) + "_" + strconv.Itoa(service.GroupID)

	if _, ok := ctxMap.Load(key); ok {
		error_nt.SendInfo(fmt.Sprintf("Service %s already exists.", key))
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	ctxMap.Store(key, cancel)

	go NotificationService(ctx, service)

	//err = client.SendNotificationMsg(
	//	"group",
	//	service.GroupID,
	//	fmt.Sprintf(
	//		"New live service has been created,RoomID:%d\n,GroupID:%d\n",
	//		service.RoomID,
	//		service.GroupID,
	//	),
	//)
	//if err != nil {
	//	log.Println(err)
	//	error_nt.SendInfo(err.Error())
	//}

	err = client.SendNotificationMsg(
		"private",
		config.Srv.Bot.QQ,
		fmt.Sprintf(
			"New live service has been created,RoomID:%d\n,GroupID:%d\n",
			service.RoomID,
			service.GroupID,
		),
	)
	if err != nil {
		log.Println(err)
		error_nt.SendInfo(err.Error())
	}
}

func NewSpaceServiceAdd(service types.BilibiliService) {
	var err error
	key := strconv.Itoa(service.UserID) + "_" + strconv.Itoa(service.GroupID)

	if _, ok := ctxMap.Load(key); ok {
		error_nt.SendInfo(fmt.Sprintf("Service %s already exists.", key))
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	ctxMap.Store(key, cancel)

	go Space(ctx, service.UserID, service.GroupID)

	//err = client.SendNotificationMsg(
	//	"group",
	//	service.GroupID,
	//	fmt.Sprintf(
	//		"New space service has been created,UserID:%d\n,GroupID:%d\n",
	//		service.UserID,
	//		service.GroupID,
	//	),
	//)
	//if err != nil {
	//	log.Println(err)
	//	error_nt.SendInfo(err.Error())
	//}

	err = client.SendNotificationMsg(
		"private",
		config.Srv.Bot.QQ,
		fmt.Sprintf(
			"New space service has been created,UserID:%d\n,GroupID:%d\n",
			service.UserID,
			service.GroupID,
		),
	)
	if err != nil {
		log.Println(err)
		error_nt.SendInfo(err.Error())
	}
}

func CancelService(ID, groupID int) {
	key := strconv.Itoa(ID) + "_" + strconv.Itoa(groupID)
	v, ok := ctxMap.LoadAndDelete(key)

	if ok {
		v.(context.CancelFunc)()
	}
}
