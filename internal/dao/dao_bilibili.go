package dao

import (
	"QQ_bot/error_nt"
	"QQ_bot/types"
	"QQ_bot/types/table"
	"github.com/gogf/gf/v2/frame/g"
	"log"
	"strconv"
)

func GetBilibiliServices() *[]types.BilibiliService {
	var (
		err      error
		db       = g.DB()
		services = make([]types.BilibiliService, 0)
	)

	m := db.Model(table.BilibiliService)
	err = m.Scan(&services)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &services
}

func UpdateBilibiliServiceInfo(serviceInfo types.BilibiliService) {
	var (
		err error
		db  = g.DB()
	)

	m := db.Model(table.BilibiliService)
	_, err = m.Replace(g.Map{
		"user_group_id":      strconv.Itoa(serviceInfo.UserID) + "_" + strconv.Itoa(serviceInfo.GroupID),
		"name":               serviceInfo.Name,
		"user_id":            serviceInfo.UserID,
		"room_id":            serviceInfo.RoomID,
		"group_id":           serviceInfo.GroupID,
		"space_notification": serviceInfo.SpaceNotification,
		"live_notification":  serviceInfo.LiveNotification,
		"at_all":             serviceInfo.AtAll,
	})
	if err != nil {
		error_nt.SendInfo(err.Error())
	}
}

func GetServiceInfoByGroupID(groupID int) *types.BilibiliService {
	var (
		err      error
		db       = g.DB()
		services = make([]types.BilibiliService, 0)
	)

	m := db.Model(table.BilibiliService)
	m.Where("group_id=", groupID)
	err = m.Scan(&services)
	if err != nil {
		log.Println(err)
		return nil
	}

	if len(services) < 1 {
		return nil
	}

	return &services[0]
}
