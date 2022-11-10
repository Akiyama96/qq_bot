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

	msg = msg[22:]

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

	err = SendNotificationMsg(msgType, id, string(data))
	if err != nil {
		log.Println(err)
	}
}
