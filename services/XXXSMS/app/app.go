package app

import (
	"flag"
	"time"

	"github.com/robert-pkg/XXX4Go/common/appbase"
	"github.com/robert-pkg/XXX4Go/common/registry"
	"github.com/robert-pkg/XXX4Go/services/XXXSMS/conf"
	"github.com/robert-pkg/XXX4Go/services/XXXSMS/service"
	"github.com/robert-pkg/micro-go/log"
	zap_log "github.com/robert-pkg/micro-go/log/zap-log"
	grpc_server "github.com/robert-pkg/micro-go/rpc/server/grpc"
)

// NewApp .
func NewApp() appbase.Application {
	return &app{}
}

// app .
type app struct {
	grpcSvr *grpc_server.Server
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

	registry := registry.InitRegistry(nil)
	app.grpcSvr = app.startGrpcServer(registry, "community.service.XXXSMS", service.New(conf.Conf))

	appbase.WaitForQuit(app)

}

func (app *app) OnQuit() {
	if app.grpcSvr != nil {
		app.grpcSvr.Shutdown() // 关闭grpc
	}

	time.Sleep(time.Second * 2)
	log.Info("exit.")
}
