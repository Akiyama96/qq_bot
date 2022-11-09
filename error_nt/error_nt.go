package error_nt

import (
	"QQ_bot/config"
	"QQ_bot/internal/client"
	"fmt"
)

func SendInfo(err string) {
	_ = client.SendNotificationMsg(
		"private",
		config.Srv.Bot.QQ,
		fmt.Sprintf(
			"Info:%s",
			err,
		),
	)
}
