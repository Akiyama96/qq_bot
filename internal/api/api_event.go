package api

import (
	"QQ_bot/internal/iservice"
	"QQ_bot/types"
	"encoding/json"
	"github.com/gogf/gf/v2/net/ghttp"
	"io"
	"log"
)

func Event(g *ghttp.RouterGroup) {
	g.POST("/", event)
}

func event(req *ghttp.Request) {
	data, err := io.ReadAll(req.Body)

	m := types.Event{}

	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Println(err)
		return
	}

	iservice.HandleEvent(&m)
}
