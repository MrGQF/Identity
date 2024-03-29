// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"gitee.com/QunXiongZhuLu/KongMing/http"
	"github.com/google/wire"
	"identify/app/dao"
	http2 "identify/app/server/http"
	"identify/app/service"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	redis, cleanup, err := dao.NewRedis()
	if err != nil {
		return nil, nil, err
	}
	daoImpl, cleanup2, err := dao.New(redis)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	client, cleanup3, err := http.NewClient()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceService, cleanup4, err := service.NewService(daoImpl, client)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	engine, cleanup5, err := http2.New(serviceService)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, cleanup6, err := NewApp(serviceService, engine)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

var DaoProvider = wire.NewSet(dao.New, dao.NewRedis)

var ServiceProvider = wire.NewSet(service.NewService, DaoProvider, http.NewClient)

var ServerProvider = wire.NewSet(http2.New, ServiceProvider)
