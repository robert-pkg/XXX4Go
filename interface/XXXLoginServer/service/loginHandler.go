package service

import (
	"context"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/robert-pkg/XXX4Go/common/ecode"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/api"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/dao"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/model"
	"github.com/robert-pkg/micro-go/log"
	"github.com/robert-pkg/micro-go/utils"
)

type loginHandler struct {
	dao *dao.Dao
}

// Login 登录
func (h *loginHandler) Handle(ctx context.Context, in *api.LoginReq) (resp *api.LoginReply, err error) {

	log.Info("enter", "name", "Login")

	resp = &api.LoginReply{
		Code:    ecode.OK.Code(),
		Message: ecode.OK.Msg(),
		Data:    nil,
	}

	deviceType, err := model.GetDeviceType(in.DeviceType)
	if err != nil {
		resp.Code = ecode.ErrDeviceType.Code()
		resp.Message = ecode.ErrDeviceType.Msg()
		return resp, nil
	}

	if isValid, err := h.checkVcode(in.Mobile, in.VerifyCode); err != nil {
		return nil, err
	} else if !isValid {
		resp.Code = eCodeLoginVcodeInvalid.Code()
		resp.Message = eCodeLoginVcodeInvalid.Msg()
		return resp, nil
	}

	userID, err := h.dao.GetUserID(in.Mobile, true)
	if err != nil {
		return nil, err
	}

	newToken, err := h.generateToken(in)
	if err != nil {
		return nil, err
	}

	expireTimeStamp := time.Now().Unix() + 60*60 // 一小时有效期
	if err := h.dao.SaveUserToken(userID, deviceType, newToken, expireTimeStamp); err != nil {
		return nil, err
	}

	h.dao.DeleteUserTokenByRedis(userID, deviceType, newToken)

	resp.Data = &api.LoginReplyData{
		UserId:   userID,
		Token:    newToken,
		ExpireTs: expireTimeStamp,
	}

	log.Info("leave", "name", "Login")

	return
}

// 生成token
func (h *loginHandler) generateToken(in *api.LoginReq) (newToken string, err error) {
	var b strings.Builder
	b.WriteString(in.Mobile)
	b.WriteString(in.DeviceType)
	b.WriteString(strconv.FormatInt(time.Now().UnixNano(), 10))
	b.WriteString(strconv.Itoa(rand.Intn(1000)))
	b.WriteString("b5b62483-e273-457c-a6d6-1b183742cf29")
	newToken = b.String()

	newToken, err = utils.GetMd5(b.String())
	if err != nil {
		log.Info("err", "err", err)
	}

	return
}

// 根据 手机号、验证码，验证合法性
func (h *loginHandler) checkVcode(mobile string, code string) (isValid bool, err error) {

	isExist, expireTimeStamp, err := h.dao.GetVerifyCodeFromRedis(mobile, code)
	if err != nil {
		return false, err
	}

	if !isExist {
		return false, nil
	}

	if expireTimeStamp <= time.Now().Unix() {
		// 过期了
		return false, nil
	}

	return true, nil
}
