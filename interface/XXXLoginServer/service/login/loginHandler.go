package login

import (
	"context"
	"math/rand"
	"strconv"
	"strings"
	"time"

	user_center_api "github.com/robert-pkg/XXX4Go/services/UserCenter/api"

	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/api"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/conf"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/dao"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/model"
	"github.com/robert-pkg/micro-go/ecode"
	"github.com/robert-pkg/micro-go/log"
	grpc_client "github.com/robert-pkg/micro-go/rpc/client/grpc"
	"github.com/robert-pkg/micro-go/utils"
)

const (
	// TokenValidSeconds token有效秒数
	TokenValidSeconds = 60 * 60 * 24 * 3 // 3天
)

type loginHandler struct {
	dao              *dao.Dao
	userCenterClient *grpc_client.Client
}

// Login 登录
func (h *loginHandler) Handle(ctx context.Context, in *api.LoginReq) (resp *api.LoginReply, err error) {

	resp = &api.LoginReply{
		Code:    ecode.OK.Code(),
		Message: ecode.OK.Message(),
		Data:    nil,
	}

	deviceType := model.GetDeviceType(in.DeviceType)
	if deviceType <= 0 {
		resp.Code, resp.Message = ecode.ErrDeviceType.CodeAndMessage()
		return resp, nil
	}

	if isValid, err := h.checkVcode(in.Mobile, in.VerifyCode); err != nil {
		return nil, err
	} else if !isValid {
		resp.Code, resp.Message = conf.LoginVcodeInvalid.CodeAndMessage()
		return resp, nil
	}

	userID, err := h.getUserIDByMobile(ctx, in.Mobile)
	if err != nil {
		return nil, err
	}

	newToken, err := h.generateToken(in)
	if err != nil {
		return nil, err
	}

	expireTimeStamp := time.Now().Unix() + TokenValidSeconds //
	if err := h.dao.SaveUserToken(userID, deviceType, newToken, expireTimeStamp); err != nil {
		return nil, err
	}

	h.dao.DeleteUserTokenByRedis(userID, deviceType, newToken)

	resp.Data = &api.LoginReplyData{
		UserId:   userID,
		Token:    newToken,
		ExpireTs: expireTimeStamp,
	}

	return
}

// 根据手机号获取 用户ID， 若用户不存在，则自动新建
func (h *loginHandler) getUserIDByMobile(ctx context.Context, mobile string) (int64, error) {

	var req user_center_api.GetUserIDByMobileReq
	req.Mobile = mobile
	req.AutoCreate = true

	var reply user_center_api.GetUserIDByMobileReply

	if err := h.userCenterClient.Call(ctx, "GetUserIDByMobile", req, &reply); err != nil {
		log.Error("err", "err", err)
		return 0, err
	}

	return reply.UserId, nil
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
