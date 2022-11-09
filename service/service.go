package service

import (
	"QQ_bot/internal/dao"
	"QQ_bot/internal/iservice"
	"fmt"
)

// LoadServiceInfo 加载服务
func LoadServiceInfo() {
	serviceInfo := dao.GetBilibiliServices()
	if serviceInfo == nil {
		fmt.Println("load service info failed or service is nil")
		return
	}

	for _, service := range *serviceInfo {

		if service.LiveNotification == 1 {
			iservice.NewLiveServiceAdd(service)
		}

		if service.SpaceNotification == 1 {
			iservice.NewSpaceServiceAdd(service)
		}
	}
}
