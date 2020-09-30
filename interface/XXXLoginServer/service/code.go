package service

import "github.com/robert-pkg/XXX4Go/common/ecode"

// 本服务定义的错误码
var (
	eCodeMobileInvalid     = ecode.New(20001, "手机号码不正确")
	eCodeLoginVcodeInvalid = ecode.New(20002, "验证码无效")
)
