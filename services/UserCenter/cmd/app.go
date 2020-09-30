package main

import (
	"flag"
	"time"

	"github.com/robert-pkg/XXX4Go/services/UserCenter/conf"
	"github.com/robert-pkg/XXX4Go/services/UserCenter/dao"
	"github.com/robert-pkg/XXX4Go/services/UserCenter/service"
	"github.com/robert-pkg/micro-go/appbase"
	"github.com/robert-pkg/micro-go/log"
	zap_log "github.com/robert-pkg/micro-go/log/zap-log"
	consul_registry "github.com/robert-pkg/micro-go/registry/consul"
	grpc_server "github.com/robert-pkg/micro-go/rpc/server/grpc"
	jaeger_trace "github.com/robert-pkg/micro-go/trace/jaeger-trace"
)

const (
	ServiceName = "community.service.UserCenter"
)

func newApp() appbase.Application {
	return &app{}
}

// app .
type app struct {
	grpcSvr *grpc_server.Server
	dao     *dao.Dao
}

// Init init App
func (app *app) Init() {

	flag.Parse()

	var err error
	// 初始化配置信息
	if err = conf.Init(); err != nil {
		panic(err)
	}

	if err = zap_log.InitByConfig(&conf.Conf.Log); err != nil {
		panic(err)
	}

	log.Info("start")

	if app.dao, err = dao.New(conf.Conf.MysqlConfig); err != nil {
		panic(err)
	}

}

// Run .
func (app *app) Run() {

	defer func() {
		if err := recover(); err != nil {
			log.Error("crash", "err", err)
		}

		log.Close()
	}()

	registry := consul_registry.InitRegistry(nil)

	_, tracerCloser, err := jaeger_trace.NewTracer(ServiceName, &conf.Conf.TraceConfig)
	if err != nil {
		panic(err)
	}
	defer tracerCloser.Close()

	app.grpcSvr = app.startGrpcServer(registry, ServiceName, service.New(app.dao))

	appbase.WaitForQuit(app)

}

func (app *app) OnQuit() {
	if app.grpcSvr != nil {
		app.grpcSvr.Shutdown() // 关闭grpc
	}

	time.Sleep(time.Second * 2)
	log.Info("exit.")
}
