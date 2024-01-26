package http

import (
	h "gitee.com/QunXiongZhuLu/KongMing/http"
	"gitee.com/QunXiongZhuLu/KongMing/log"
	"gitee.com/QunXiongZhuLu/KongMing/util"
	"gitee.com/QunXiongZhuLu/kratos/pkg/conf/paladin"
	bm "gitee.com/QunXiongZhuLu/kratos/pkg/net/http/blademaster"
	"identify/app/service"
)

var (
	svc *service.Service
)

func New(s *service.Service) (*bm.Engine, func(), error) {
	svc = s
	engine, cf, err := NewServer()
	return engine, cf, err
}

// New new a bm server.
// func New(s pb.DemoServer) (engine *bm.Engine, err error) {
func NewServer() (engine *bm.Engine, cf func(), err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.yaml").Unmarshal(&ct); err != nil {
		log.TLog.Error("httpserver config unmarshal error\n", err)
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		log.TLog.Error("httpserver config unmarshaltoml error\n", err)
		return
	}

	// init http server
	engine = bm.DefaultServer(&cfg)
	initRouter(engine)

	// close func
	cf = func() {
	}

	return
}

// 初始化路由
func initRouter(e *bm.Engine) {

	if e == nil {
		log.TLog.Error("Create Http Server Error")
	}

	// add validate
	headValidateHandle := h.GetHeadValidateHandle()

	g := e.Group("/validate")
	{
		g.GET("/checkalive", headValidateHandle, CheckAlive)
		g.GET("/getuserinfo", headValidateHandle, GetUserInfoByToken)
		g.POST("/refreshtoken", headValidateHandle, RefreshToken) // 刷新Token
		g.POST("/logout", headValidateHandle, Logout)             // 注销登录
		g.POST("/getpassport", headValidateHandle, GetPCPassport) // 获取通行证
		g.POST("/getcookie", headValidateHandle, GetCookie)       // 获取Cookie
	}

}

func CheckAlive(c *bm.Context) {
	defer util.TimeCost("checkalive")()

	c.JSON("OK", nil)
}
