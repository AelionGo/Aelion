// Package types 请求和响应体的定义
package types

type CaptchaRequest struct {
	Type string
}

type CaptchaResponse struct {
	Enabled bool   `json:"enabled"` // 是否启用验证码
	Id      string `json:"id"`      // 验证码ID
	B64s    string `json:"b64s"`    // 验证码图片的Base64编码
}
