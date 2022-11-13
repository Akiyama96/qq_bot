package client

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"log"
)

func Xiaoai(msgType string, id int, msg string) {
	client := g.Client()

	url := fmt.Sprintf("http://81.70.100.130/api/xiaoai.php?msg=%s&n=text", msg)

	res, err := client.Get(context.Background(), url)
	if err != nil {
		log.Println(err)
	}

	if res.Response.StatusCode != 200 {
		err = SendNotificationMsg(msgType, id, "bot不会")
		return
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	msg = string(data)

	if msg == "对不起，暂不支持该功能，你和我聊点儿别的吧" {
		msg = "[CQ:image,file=8bf34b1019c1419558666ddb73a903d8.image,url=https://c2cpicdw.qpic.cn/offpic_new/1131568220//1131568220-1507055459-8BF34B1019C1419558666DDB73A903D8/0?term=3&amp;is_origin=0]"
	}

	err = SendNotificationMsg(msgType, id, msg)
	if err != nil {
		log.Println(err)
	}
}
