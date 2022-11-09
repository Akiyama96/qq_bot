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

func GetPic(tags []string, r18 int) *types.PicInfo {
	var picInfo = types.PicInfo{}
	client := g.Client()

	req := g.Map{
		"r18":         r18,
		"num":         1,
		"replace_url": "https://i.pixiv.re",
		"tags":        tags,
	}

	client.ContentJson()

	jdata, err := json.Marshal(req)
	if err != nil {
		log.Println(jdata)
	}

	res, err := client.Post(context.Background(), picUrl, jdata)
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
