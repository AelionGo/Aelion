// Package msg 构建响应
package msg

import "github.com/AelionGo/Aelion/pkg/errors"

type Response struct {
	Code int    `json:"code"` // 响应码
	Msg  string `json:"msg"`  // 响应消息
	Data any    `json:"data"` // 响应数据
}

func GetResponse(code errors.Code, data any) *Response {
	if data == nil {
		data = []struct{}{} // 如果data为nil，返回空对象
	}
	return &Response{
		Code: code,
		Msg:  errors.GetMsg(code),
		Data: data,
	}
}
