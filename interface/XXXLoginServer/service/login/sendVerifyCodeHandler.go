package login

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/api"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/conf"
	"github.com/robert-pkg/XXX4Go/interface/XXXLoginServer/dao"
	"github.com/robert-pkg/micro-go/ecode"
	"github.com/robert-pkg/micro-go/log"
	grpc_client "github.com/robert-pkg/micro-go/rpc/client/grpc"

	sms_api "github.com/robert-pkg/XXX4Go/services/XXXSMS/api"
)

type sendVerifyCodeHandler struct {
	dao       *dao.Dao
	smsClient *grpc_client.Client
}

func (h *sendVerifyCodeHandler) checkMobile(mobile string) bool {
	if len(mobile) != 11 {
		return false
	}

	return true
}

func (h *sendVerifyCodeHandler) Handle(ctx context.Context, in *api.SendVerifyCodeReq) (resp *api.SendVerifyCodeReply, err error) {

	resp = &api.SendVerifyCodeReply{}
	resp.Code, resp.Message = ecode.OK.CodeAndMessage()

	if !h.checkMobile(in.Mobile) {
		resp.Code, resp.Message = conf.MobileInvalid.CodeAndMessage()
		return resp, nil
	}

	verifyCode := h.generateVcode(4, 0)
	vCodeMsg := fmt.Sprintf("【XXX】您好，%s是您本次登录的验证码，该验证码在10分钟内有效。", verifyCode)

	resp.Vcode = verifyCode // 为了调试方便，直接返回验证码
	if err = h.sendVcodeImp(ctx, in.Mobile, vCodeMsg); err != nil {
		return nil, err
	}

	expireTimeStamp := time.Now().Unix() + 5*60
	if err = h.dao.SetVerifyCode2Redis(in.Mobile, verifyCode, expireTimeStamp); err != nil {
		return nil, err
	}

	return resp, err
}

func (h *sendVerifyCodeHandler) sendVcodeImp(ctx context.Context, mobile string, msg string) error {

	req := &sms_api.SendMsgReq{
		Mobile: mobile,
		Msg:    msg,
	}

	if err := h.smsClient.Call(ctx, "SendMsg", req, nil); err != nil {
		log.Error("err", "err", err)
		return err
	}

	return nil
}

/*
 * generateVcode 生成验证码
 * size 随机码的位数
 * kind 验证码类型, 0 纯数字 1 小写字母 2 大写字母 3 数字、大小写字母
 */
func (h *sendVerifyCodeHandler) generateVcode(size int, kind int) string {

	ikind := kind
	kinds := [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}} // 数字的ASCII码从48开始， A是65， a是97
	result := make([]byte, size)

	isAll := false
	if kind > 2 || kind < 0 {
		isAll = true
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}

		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
