package service

import (
	"context"
	"time"

	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/api"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/dao"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/model"
	"github.com/robert-pkg/micro-go/log"
)

type verifyTokenHandler struct {
	dao *dao.Dao
}

func (h *verifyTokenHandler) Handle(ctx context.Context, in *api.VerifyTokenReq) (resp *api.VerifyTokenReply, err error) {
	log.Info("enter", "name", "VerifyToken")

	resp = &api.VerifyTokenReply{
		IsValid:  false,
		ExpireTs: 0,
	}

	deviceType, err := model.GetDeviceType(in.DeviceType)
	if err != nil {
		// 设备类型不正确，一律当token无效处理
		return resp, nil
	}

	// 检查redis
	if isExist, expireTimeStamp, err := h.dao.GetUserTokenFromRedis(in.UserId, deviceType, in.Token); err != nil {
		return nil, err
	} else if isExist {
		if expireTimeStamp <= time.Now().Unix() {
			return resp, nil // 过期了， 2无效token
		}

		resp.IsValid = true
		resp.ExpireTs = expireTimeStamp
		return resp, nil
	}

	// 从db中检查
	if isExist, expireTimeStamp, err := h.dao.GetUserToken(in.UserId, deviceType, in.Token); err != nil {
		return nil, err
	} else if isExist {
		if expireTimeStamp <= time.Now().Unix() {
			resp.IsValid = false // 过期了， 2无效token
		} else {
			resp.IsValid = true
			resp.ExpireTs = expireTimeStamp
		}
	} else {
		resp.IsValid = false // 不存在 2无效token
	}

	if resp.IsValid {
		// 保存到redis
		h.dao.SetUserToken2Redis(in.UserId, deviceType, in.Token, resp.ExpireTs, false)
	} else {
		// 插入一个 过期token，防无效token的流量攻击
		h.dao.SetUserToken2Redis(in.UserId, deviceType, in.Token, 0, true)
	}

	return resp, nil

}
