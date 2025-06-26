// Filename: routes.go
// Description: HTTP路由

// Package router 定义路由和处理请求
package router

import (
	"github.com/AelionGo/Aelion/internal/router/handler/v1/ping"
	userx "github.com/AelionGo/Aelion/internal/router/handler/v1/user"
	"github.com/AelionGo/Aelion/internal/svc"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRoutes(s *server.Hertz, svcCtx *svc.ServiceContext) {
	v1 := s.Group("/v1")
	v1.GET("/ping", ping.PingHandler(svcCtx)) // 健康检查

	user := v1.Group("/user")
	user.GET("/captcha", userx.CaptchaHandler(svcCtx))    // 获取图形验证码
	user.POST("/register", userx.RegisterHandler(svcCtx)) // 用户注册
}
