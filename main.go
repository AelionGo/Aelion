// Package main

// Filename: main.go
// Description: 项目入口

package main

import (
	"flag"
	"fmt"
	"github.com/AelionGo/Aelion/internal/config"
	"github.com/AelionGo/Aelion/internal/router"
	"github.com/AelionGo/Aelion/internal/svc"
	"github.com/AelionGo/Aelion/models"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	// 解析命令行参数
	flag.Parse()

	//初始化数据库连接
	err := models.InitDB()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize database: %v", err))
	}

	// 加载配置文件
	c, err := config.Init(*configFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// 创建HTTP服务器
	server := router.NewServer(&c.Server)

	// 创建服务上下文
	svcCtx := svc.NewServiceContext(c)

	// 注册路由
	router.RegisterRoutes(server, svcCtx)

	// 启动服务器
	fmt.Printf("Starting server at %s:%d...\n", c.Server.Host, c.Server.Port)
	server.Spin()
}
