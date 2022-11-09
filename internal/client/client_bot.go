package client

import (
	"QQ_bot/config"
	"QQ_bot/types"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
)

const (
	sendMsgUrl     = "/send_msg"
	reqGroupAddUrl = "/set_group_add_request"
)

func SendNotificationMsg(messageType string, id int, msg string) error {
	serverUrl := config.Srv.Bot.Address + ":" + strconv.Itoa(config.Srv.Bot.Port)
	client := g.Client()

	senMsg := g.Map{
		"message_type": messageType,
		"message":      msg,
		"auto_escape":  false,
	}

	switch messageType {
	case "group":
		senMsg["group_id"] = id
	case "private":
		senMsg["user_id"] = id
	}

	jsonData, err := json.Marshal(senMsg)
	if err != nil {
		return err
	}

	_, err = client.Post(context.Background(), serverUrl+sendMsgUrl, jsonData)
	if err != nil {
		return err
	}

	return nil
}

func AppendInvite(data *types.Event) error {
	serverUrl := config.Srv.Bot.Address + ":" + strconv.Itoa(config.Srv.Bot.Port)
	client := g.Client()

	senMsg := g.Map{
		"flag":     data.Flag,
		"sub_type": data.SubType,
		"approve":  true,
	}

	jsonData, err := json.Marshal(senMsg)
	if err != nil {
		return err
	}

	_, err = client.Post(context.Background(), serverUrl+reqGroupAddUrl, jsonData)
	if err != nil {
		return err
	}

	return nil
}
