package types

type RegisterRequest struct {
	CaptchaId     string `json:"captcha_id"`     // 验证码ID
	CaptchaAnswer string `json:"captcha_answer"` // 验证码答案
	Nickname      string `json:"nickname"`       // 昵称
	Email         string `json:"email"`          // 邮箱
	Phone         string `json:"phone"`          // 手机号
	Password      string `json:"password"`       // 密码
	Avatar        string `json:"avatar"`         // 头像
	Group         string `json:"group"`          // 用户组
}

type RegisterResponse struct {
	Id    string `json:"id"`    // 用户ID
	Token string `json:"token"` // JWT令牌
}
