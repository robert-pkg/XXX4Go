package ecode

// common ecode
var (
	// OK 成功
	OK = add(0, "OK")
	// ErrSystemBusy .
	ErrSystemBusy = add(50001, "服务繁忙，请稍后重试")
	// ErrDeviceType .
	ErrDeviceType = add(50002, "设备类型错误")
	// ErrNoUserID .
	ErrNoUserID = add(50003, "用户ID缺失")
	// ErrParam .
	ErrParam = add(50004, "输入参数错误")
)

// 各服务可以定义的错误码范围：20000-40000
