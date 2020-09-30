package main

import (
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
)

func main() {
	// 发送短信
	client := ypclnt.New("apikey")
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = "18616020610"
	param[ypclnt.TEXT] = "【云片网】您的验证码是1234"

	r := client.Sms().SingleSend(param)

	if r.Code != 0 {

	}
}
