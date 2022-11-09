package router

import (
	"QQ_bot/config"
	"QQ_bot/internal/api"
	"github.com/gogf/gf/v2/frame/g"
)

func SignUp() {
	//新建服务对象
	var s = g.Server("api")
	//设置服务端口
	s.SetPort(config.Srv.Http.Port)
	//设置路由组
	group := s.Group("/api")
	api.Event(group)

	s.Run()
}
