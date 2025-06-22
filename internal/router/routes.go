// package router 定义路由和处理请求

// Filename: routes.go
// Description: HTTP路由

package router

import (
	"github.com/AelionGo/Aelion/internal/router/handler/v1/ping"
	"github.com/AelionGo/Aelion/internal/svc"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRoutes(s *server.Hertz, svcCtx *svc.ServiceContext) {
	v1 := s.Group("/v1")

	v1.GET("/ping", ping.PingHandler(svcCtx)) // 健康检查

}
