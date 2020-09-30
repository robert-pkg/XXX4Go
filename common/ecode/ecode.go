package ecode

import (
	"fmt"
	"strconv"
)

var (
	_codeMap = map[int32]string{} // register codes.
)

// New new a ecode.Codes by int value.
// NOTE: ecode must unique in global, the New will check repeat and then panic.
// business ecode must between 20000-40000
func New(e int32, msg string) Code {

	if e < 20000 || e > 40000 {
		panic("business ecode must between [20000,40000]")
	}
	return add(e, msg)
}

func add(e int32, msg string) Code {
	if _, ok := _codeMap[e]; ok {
		panic(fmt.Sprintf("wcode: %d already exist", e))
	}
	_codeMap[e] = msg
	return Code{code: e, msg: msg}
}

// A Code is an int error code spec.
type Code struct {
	code int32
	msg  string
}

func (e Code) Error() string {
	return strconv.FormatInt(int64(e.code), 10)
}

// Code return error code
func (e Code) Code() int32 { return e.code }

// Msg return error message
func (e Code) Msg(args ...interface{}) string {
	if e.msg != "" {
		return fmt.Sprintf(e.msg, args...)
	}

	return e.Error()
}
