package types

type LoginRequest struct {
	CaptchaId     string `json:"captcha_id"`      // 验证码ID
	CaptchaAnswer string `json:"captcha_answer"`  // 验证码答案
	Email         string `json:"email,omitempty"` // 邮箱
	Phone         string `json:"phone,omitempty"` // 手机号
	Password      string `json:"password"`        // 密码
}

type LoginResponse struct {
	Id    string `json:"id"`    // 用户ID
	Token string `json:"token"` // JWT令牌
}
