package service

import (
	"context"
	"errors"

	"github.com/robert-pkg/XXX4Go/services/UserCenter/api"
	"github.com/robert-pkg/XXX4Go/services/UserCenter/dao"
)

// New init
func New(dao *dao.Dao) *Service {

	s := &Service{
		dao: dao,
	}

	return s
}

// Service 实现pb中接口
type Service struct {
	dao *dao.Dao
}

// GetUserIDByMobile .
func (s *Service) GetUserIDByMobile(ctx context.Context, in *api.GetUserIDByMobileReq) (resp *api.GetUserIDByMobileReply, err error) {

	resp = &api.GetUserIDByMobileReply{}

	if len(in.Mobile) <= 0 {
		return nil, errors.New("参数错误，手机号不能为空")
	}

	resp.UserId, resp.IsNew, err = s.dao.GetUserID(in.Mobile, in.AutoCreate)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
