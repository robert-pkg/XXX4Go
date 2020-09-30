package conf

import (
	"github.com/robert-pkg/micro-go/ecode"
)

var (
	MobileInvalidCode     int32 = 20001
	LoginVcodeInvalidCode int32 = 20002
)

// 本服务定义的错误码
var (
	MobileInvalid     = ecode.New(MobileInvalidCode)
	LoginVcodeInvalid = ecode.New(LoginVcodeInvalidCode)
)

var codeMap = map[int32]string{
	MobileInvalidCode:     "手机号码不正确",
	LoginVcodeInvalidCode: "验证码无效",
}

func registerCodes() {
	ecode.Register(codeMap)
}
