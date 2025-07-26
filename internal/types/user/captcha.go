package user

type CaptchaRequest struct {
	Type string
}

type CaptchaResponse struct {
	Enabled bool   `json:"enabled"` // 是否启用验证码
	Id      string `json:"id"`      // 验证码ID
	B64s    string `json:"b64s"`    // 验证码图片的Base64编码
}
