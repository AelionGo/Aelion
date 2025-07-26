// Package errors 定义错误码和错误信息
package errors

type Code = int

const (
	OK = 0

	ParamsError        = 40000
	CaptchaVerifyError = 40001
	EmailExists        = 40002
	PhoneExists        = 40003
	UserNotFound       = 40004
	PasswordError      = 40005
	PermissionDenied   = 40006

	GetConfigItemError   = 50000
	CaptchaGenerateError = 50001
	DatabaseError        = 50002
	TokenGenerateError   = 50003
	PasswordHashError    = 50004
)

var errMsg = map[int]string{
	OK: "",

	ParamsError:        "request params error",
	CaptchaVerifyError: "captcha verify error",
	EmailExists:        "email already exists",
	PhoneExists:        "phone already exists",
	UserNotFound:       "user not found",
	PasswordError:      "password error",
	PermissionDenied:   "permission denied",

	GetConfigItemError:   "get config item error",
	CaptchaGenerateError: "captcha generate error",
	DatabaseError:        "database error",
	TokenGenerateError:   "token generate error",
	PasswordHashError:    "password hash error",
}

func GetMsg(code int) string {
	if msg, ok := errMsg[code]; ok {
		return msg
	}
	return "unknown error"
}
