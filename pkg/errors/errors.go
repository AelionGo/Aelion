// Package errors 定义错误码和错误信息
package errors

type Code = int

const (
	OK          = 0
	ParamsError = 40000
)

var errMsg = map[int]string{
	OK:          "",
	ParamsError: "request params error",
}

func GetMsg(code int) string {
	if msg, ok := errMsg[code]; ok {
		return msg
	}
	return "unknown error"
}
