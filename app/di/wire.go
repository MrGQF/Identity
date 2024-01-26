// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	km "gitee.com/QunXiongZhuLu/KongMing/http"
	"github.com/google/wire"
	"identify/app/dao"
	"identify/app/server/http"
	"identify/app/service"
)

var DaoProvider = wire.NewSet(dao.New, dao.NewRedis)
var ServiceProvider = wire.NewSet(service.NewService, DaoProvider, km.NewClient)
var ServerProvider = wire.NewSet(http.New, ServiceProvider)

func InitApp() (*App, func(), error) {
	panic(wire.Build(ServerProvider, NewApp))
}
