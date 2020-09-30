package login

import (
	"context"

	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/api"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/dao"
	grpc_client "github.com/robert-pkg/micro-go/rpc/client/grpc"
)

// New init
func New(dao *dao.Dao, smsClient *grpc_client.Client, userCenterClient *grpc_client.Client) *Service {

	s := &Service{
		smsClient:        smsClient,
		userCenterClient: userCenterClient,
		dao:              dao,
	}

	return s
}

// Service 实现pb中接口
type Service struct {
	smsClient        *grpc_client.Client
	userCenterClient *grpc_client.Client
	dao              *dao.Dao
}

// SendVerifyCode 发送验证码
func (s *Service) SendVerifyCode(ctx context.Context, in *api.SendVerifyCodeReq) (resp *api.SendVerifyCodeReply, err error) {
	//log.Info("SendVerifyCode")

	h := &sendVerifyCodeHandler{
		dao:       s.dao,
		smsClient: s.smsClient,
	}
	return h.Handle(ctx, in)
}

// Login 登录
func (s *Service) Login(ctx context.Context, in *api.LoginReq) (resp *api.LoginReply, err error) {
	h := &loginHandler{
		dao:              s.dao,
		userCenterClient: s.userCenterClient,
	}

	return h.Handle(ctx, in)
}

// VerifyToken verify token
func (s *Service) VerifyToken(ctx context.Context, in *api.VerifyTokenReq) (resp *api.VerifyTokenReply, err error) {

	h := &verifyTokenHandler{
		dao: s.dao,
	}

	return h.Handle(ctx, in)

}
