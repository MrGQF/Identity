package main

import (
	"gitee.com/QunXiongZhuLu/KongMing/config"
	"gitee.com/QunXiongZhuLu/KongMing/log"
	"github.com/TarsCloud/TarsGo/tars"
	_ "github.com/TarsCloud/TarsGo/tars"
	"identify/app/di"
)

func main() {
	var (
		err error
		app *di.App
	)

	// Init Config
	var conf = config.Config{
		Type: config.Apollo,
		ApolloConfig: config.ApolloConfig{
			AppID:      config.IdentifyServer,
			Cluster:    config.Default,
			MetaAddr:   "http://security-config-apolloconfigservice:8080",
			Namespaces: "application.yaml,http.yaml,redis.yaml",
		},
	}
	if err = config.Init(conf); err != nil {
		log.TLog.Error("Init Config err:", err)
		panic(err)
	}

	// Get server config
	cfg := tars.GetServerConfig()

	// Init App
	app, _, err = di.InitApp()
	if err != nil {
		panic(err)
	}

	// Register Http Servant
	mux := &tars.TarsHttpMux{}
	mux.Handle("/", app.Http)
	tars.AddHttpServant(mux, cfg.App+"."+cfg.Server+".validateObj")

	// Run application
	tars.Run()
}
