package app

import (
	"flag"
	"math/rand"
	"time"

	"github.com/robert-pkg/XXX4Go/common/appbase"
	"github.com/robert-pkg/XXX4Go/common/registry"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/conf"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/dao"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/service"
	"github.com/robert-pkg/micro-go/log"
	zap_log "github.com/robert-pkg/micro-go/log/zap-log"
	grpc_client "github.com/robert-pkg/micro-go/rpc/client/grpc"
	grpc_server "github.com/robert-pkg/micro-go/rpc/server/grpc"
)

// NewApp .
func NewApp() appbase.Application {
	return &app{}
}

// app .
type app struct {
	dao     *dao.Dao
	grpcSvr *grpc_server.Server
}

// Init init App
func (app *app) Init() {

	var err error

	flag.Parse()

	// 初始化配置信息
	if err = conf.Init(); err != nil {
		panic(err)
	}

	if err = zap_log.InitByConfig(&conf.Conf.Log); err != nil {
		panic(err)
	}

	log.Info("start")

	rand.Seed(time.Now().UnixNano())
	if app.dao, err = dao.New(conf.Conf); err != nil {
		panic(err)
	}
}

func (app *app) initSMSClient() *grpc_client.Client {
	c, err := grpc_client.NewClient("community.service.XXXSMS")
	if err != nil {
		panic(err)
	}

	return c
}

// Run .
func (app *app) Run() {

	defer func() {
		if err := recover(); err != nil {
			log.Error("crash", "err", err)
		}

		log.Close()
	}()

	registry := registry.InitRegistryAsDefault(nil)
	smsClient := app.initSMSClient()

	app.grpcSvr = app.startGrpcServer(registry, "community.interface.XXXLoginServer", service.New(app.dao, smsClient))

	appbase.WaitForQuit(app)
}

func (app *app) OnQuit() {
	if app.grpcSvr != nil {
		app.grpcSvr.Shutdown() // 关闭grpc
	}

	time.Sleep(time.Second * 2)
	log.Info("exit.")
}
