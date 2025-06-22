// package router 定义路由和处理请求

// Filename: server.go
// Description: HTTP服务器

package router

import (
	"github.com/AelionGo/Aelion/internal/config"
	"github.com/cloudwego/hertz/pkg/app/server"
	"strconv"
	"strings"
)

func NewServer(c config.ServerConfig) *server.Hertz {
	port := strconv.Itoa(c.Port)
	return server.Default(server.WithHostPorts(strings.Join([]string{c.Host, port}, ":")), server.WithDisablePrintRoute(true))
}
