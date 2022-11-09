package main

import (
	"QQ_bot/log"
	"QQ_bot/service"
	"QQ_bot/service/router"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	log.InitLog()

	go router.SignUp()

	go service.LoadServiceInfo()

	g.Wait()
}
