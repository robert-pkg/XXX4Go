package service

import (
	"context"

	"github.com/robert-pkg/XXX4Go/interface/XXXUserServer/api"
	"github.com/robert-pkg/XXX4Go/interface/XXXUserServer/conf"
	"github.com/robert-pkg/XXX4Go/interface/XXXUserServer/dao"
	"github.com/robert-pkg/micro-go/log"
	grpc_server "github.com/robert-pkg/micro-go/rpc/server/grpc"
)

// New init
func New(c *conf.Config) *Service {

	dao, err := dao.New(c)
	if err != nil {
		panic(err)
	}

	s := &Service{
		c:   c,
		dao: dao,
	}

	return s
}

// Service 实现pb中接口
type Service struct {
	c *conf.Config

	dao *dao.Dao
}

// SayHello .
func (s *Service) SayHello(ctx context.Context, in *api.SayHelloReq) (resp *api.SayHelloReply, err error) {

	userID := grpc_server.GetUserIDFromCtx(ctx)
	log.Info("enter", "SayHello", "SayHello", "userID", userID)

	resp = &api.SayHelloReply{
		Reply: "Hello " + in.Message,
	}

	log.Info("leave", "SayHello", "SayHello")

	return
}
