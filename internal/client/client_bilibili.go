package client

import (
	"QQ_bot/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"log"
)

const (
	getLiveRoomInfo = "https://api.live.bilibili.com/room/v1/Room/get_info?room_id="
	getStatInfo     = "https://api.bilibili.com/x/relation/stat?vmid="
	getSpaceInfo    = "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space?offset=&host_mid="
)

func GetLiveRoomInfo(roomId string) *types.LiveRoomInfo {
	var liveRoomInfo types.LiveRoomInfo
	client := g.Client()

	res, err := client.Get(context.Background(), getLiveRoomInfo+roomId)
	if err != nil {
		log.Println(err)
		return nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = json.Unmarshal(data, &liveRoomInfo)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &liveRoomInfo
}

func GetUserStatInfo(vmid string) *types.UserStatInfo {
	var userStatInfo = types.UserStatInfo{}
	client := g.Client()

	res, err := client.Get(context.Background(), getStatInfo+vmid+"&jsonp=jsonp")
	if err != nil {
		log.Println(err)
		return nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = json.Unmarshal(data, &userStatInfo)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &userStatInfo
}

func GetSpaceInfo(mid string) *types.SpaceInfo {
	var SpaceInfo = types.SpaceInfo{}

	client := g.Client()

	res, err := client.Get(context.Background(), getSpaceInfo+mid+"&timezone_offset=-480")
	if err != nil {
		log.Println(err)
		return nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = json.Unmarshal(data, &SpaceInfo)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &SpaceInfo
}

func GetCardInfo(uid int) *types.BilibiliUserInfo {
	var SpaceInfo = types.BilibiliUserInfo{}

	client := g.Client()

	res, err := client.Get(context.Background(), fmt.Sprintf("https://api.bilibili.com/x/web-interface/card?mid=%d", uid))
	if err != nil {
		log.Println(err)
		return nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = json.Unmarshal(data, &SpaceInfo)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &SpaceInfo
}
