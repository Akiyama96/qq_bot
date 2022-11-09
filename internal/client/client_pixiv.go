package client

import (
	"QQ_bot/types"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"log"
)

const picUrl = "https://setu.yuban10703.xyz/setu"

func GetPic(tags []string) *types.PicInfo {
	var picInfo = types.PicInfo{}
	client := g.Client()

	req := g.Map{
		"r18":         1,
		"num":         1,
		"tags":        tags,
		"replace_url": "https://i.pixiv.re",
	}

	jdata, err := json.Marshal(req)
	if err != nil {
		log.Println(jdata)
	}

	res, err := client.Get(context.Background(), picUrl, jdata)
	if err != nil {
		log.Println(err)
		return nil
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = json.Unmarshal(data, &picInfo)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &picInfo
}
