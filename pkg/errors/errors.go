// Package errors 定义错误码和错误信息
package errors

type Code = int

const (
	OK = 0

	ParamsError = 40000

	GetConfigItemError   = 50000
	CaptchaGenerateError = 50001
)

var errMsg = map[int]string{
	OK: "",

	ParamsError: "request params error",

	GetConfigItemError:   "get config item error",
	CaptchaGenerateError: "captcha generate error",
}

func GetMsg(code int) string {
	if msg, ok := errMsg[code]; ok {
		return msg
	}
	return "unknown error"
}
