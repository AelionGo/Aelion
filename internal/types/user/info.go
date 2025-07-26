package user

type InfoRequest struct {
	Id string `json:"id"` // 用户ID
}

type InfoResponse struct {
	Id       string `json:"id"`       // 用户ID
	Nickname string `json:"nickname"` // 昵称
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 手机号
	Avatar   string `json:"avatar"`   // 头像
	Group    string `json:"group"`    // 用户组
}
