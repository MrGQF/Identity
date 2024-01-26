package di

import (
	"context"
	bm "gitee.com/QunXiongZhuLu/kratos/pkg/net/http/blademaster"
	"identify/app/service"
	"time"
)

//go:generate kratos tool wire
type App struct {
	Svc  *service.Service
	Http *bm.Engine
}

func NewApp(svc *service.Service, h *bm.Engine) (app *App, closeFunc func(), err error) {
	app = &App{
		Svc:  svc,
		Http: h,
	}
	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)

		if err := h.Shutdown(ctx); err != nil {
		}
		cancel()
	}
	return
}
