// Package msg 构建响应
package msg

import "github.com/AelionGo/Aelion/pkg/errors"

type Response struct {
	Code int    `json:"code"` // 响应码
	Msg  string `json:"msg"`  // 响应消息
	Data any    `json:"data"` // 响应数据
}

func GetResponse(code errors.Code, data any) *Response {
	return &Response{
		Code: code,
		Msg:  errors.GetMsg(code),
		Data: data,
	}
}
