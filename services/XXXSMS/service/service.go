package service

import (
	"context"

	"github.com/robert-pkg/XXX4Go/services/XXXSMS/api"
	"github.com/robert-pkg/XXX4Go/services/XXXSMS/conf"
	"github.com/robert-pkg/micro-go/log"
)

// New init
func New(c *conf.Config) *Service {

	s := &Service{
		c: c,
	}

	return s
}

// Service 实现pb中接口
type Service struct {
	c *conf.Config
}

// SendMsg 发送短信
func (s *Service) SendMsg(ctx context.Context, in *api.SendMsgReq) (resp *api.NoReply, err error) {

	resp = &api.NoReply{}

	log.Info("SendMsg")

	return resp, nil
}
