package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/conf"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/dao"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/http"
	"github.com/robert-pkg/micro-go/appbase"
	"github.com/robert-pkg/micro-go/log"
	zap_log "github.com/robert-pkg/micro-go/log/zap-log"
	consul_registry "github.com/robert-pkg/micro-go/registry/consul"
	jaeger_trace "github.com/robert-pkg/micro-go/trace/jaeger-trace"
)

const (
	// XXXLoginServer .
	XXXLoginServer = "community.interface.XXXLoginServer"
)

func newApp() appbase.Application {
	return &app{}
}

// app .
type app struct {
	dao *dao.Dao
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
	if app.dao, err = dao.New(conf.Conf.MysqlConfig, conf.Conf.RedisConfig); err != nil {
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

	_, tracerCloser, err := jaeger_trace.NewTracer(XXXLoginServer, &conf.Conf.TraceConfig)
	if err != nil {
		panic(err)
	}
	defer tracerCloser.Close()

	registry := consul_registry.InitRegistryAsDefault(nil)

	if err := http.Init(registry, XXXLoginServer, app.dao); err != nil {
		panic(errors.Wrap(err, "http server start fail"))
	}

	appbase.WaitForQuit(app)
}

func (app *app) OnQuit() {
	http.Close()

	time.Sleep(time.Second * 2)
	log.Info("exit.")
}
