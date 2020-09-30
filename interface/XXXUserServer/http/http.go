package http

import (
	"net/http"

	"github.com/robert-pkg/XXX4Go/interface/XXXUserServer/api"
	"github.com/robert-pkg/XXX4Go/interface/XXXUserServer/service"

	"github.com/robert-pkg/micro-go/log"
	grpc_client "github.com/robert-pkg/micro-go/rpc/client/grpc"

	"github.com/gin-gonic/gin"

	"github.com/robert-pkg/micro-go/registry"
	http_svr "github.com/robert-pkg/micro-go/rpc/server/http"
)

const (
	ServiceName_UserCenter = "community.service.UserCenter"
)

var (
	svc     *service.Service
	httpSvr *http_svr.Server

	userCenterClient *grpc_client.Client
)

// Close .
func Close() {
	if httpSvr != nil {
		httpSvr.Shutdown()
	}
}

// Init .
func Init(registry registry.Registry, serviceName string) error {

	initService()

	httpSvr = http_svr.NewServer(registry, serviceName)

	// 初始化gin
	setupRouter(httpSvr.GetEngine(), httpSvr.GetShortServiceName())

	if err := httpSvr.Start(); err != nil {
		return err
	}

	return nil
}

func initService() {

	var err error
	userCenterClient, err = grpc_client.NewClient(ServiceName_UserCenter)
	if err != nil {
		panic(err)
	}

	svc = service.New()
}

func setupRouter(engine *gin.Engine, serviceName string) {

	groupName := "/api/" + serviceName
	log.Info(groupName)
	myGroup := engine.Group(groupName)
	if true {
		h := &userHandler{}
		myGroup.POST("/GetUserName", h.GetUserName)
	}

	//没有路由的页面
	//为没有配置处理函数的路由添加处理程序，默认情况下它返回404
	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found!")
	})
}

type userHandler struct {
}

func (s *userHandler) GetUserName(c *gin.Context) {

	req := &api.NoArgRequest{}
	if err := c.BindJSON(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := http_svr.GetContext(c)
	defer cancel() // terminal routine tree

	resp, err := svc.GetUserName(ctx, req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
