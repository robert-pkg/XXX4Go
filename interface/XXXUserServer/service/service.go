package service

import (
	"context"

	"github.com/robert-pkg/XXX4Go/interface/XXXUserServer/api"
	"github.com/robert-pkg/micro-go/log"
	"github.com/robert-pkg/micro-go/rpc"
)

// New init
func New() *Service {

	s := &Service{}

	return s
}

// Service 实现pb中接口
type Service struct {
}

// GetUserName .
func (s *Service) GetUserName(ctx context.Context, in *api.NoArgRequest) (resp *api.GetUserNameReply, err error) {

	userID := rpc.GetUserIDFromCtx(ctx)

	log.Info("enter", "SayHello", "SayHello", "userID", userID)

	resp = &api.GetUserNameReply{
		Name: "Hello ",
	}

	return
}
