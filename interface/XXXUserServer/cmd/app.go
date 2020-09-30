package main

import (
	"flag"
	"time"

	"github.com/pkg/errors"
	"github.com/robert-pkg/XXX4Go/interface/XXXUserServer/conf"
	"github.com/robert-pkg/XXX4Go/interface/XXXUserServer/http"
	"github.com/robert-pkg/micro-go/appbase"
	"github.com/robert-pkg/micro-go/log"
	zap_log "github.com/robert-pkg/micro-go/log/zap-log"
	consul_registry "github.com/robert-pkg/micro-go/registry/consul"
)

const (
	ServiceName = "community.interface.XXXUserServer"
)

// newApp .
func newApp() appbase.Application {
	return &app{}
}

// app .
type app struct {
}

// Init init App
func (app *app) Init() {
	flag.Parse()

	// 初始化配置信息
	if err := conf.Init(); err != nil {
		panic(err)
	}

	if err := zap_log.InitByConfig(&conf.Conf.Log); err != nil {
		panic(err)
	}

	log.Info("start")
}

// Run .
func (app *app) Run() {

	defer func() {
		if err := recover(); err != nil {
			log.Error("crash", "err", err)
		}

		log.Close()
	}()

	registry := consul_registry.InitRegistryAsDefault(nil)

	if err := http.Init(registry, ServiceName); err != nil {
		panic(errors.Wrap(err, "http server start fail"))
	}

	appbase.WaitForQuit(app)

}

func (app *app) OnQuit() {
	http.Close()

	time.Sleep(time.Second * 2)
	log.Info("exit.")
}
