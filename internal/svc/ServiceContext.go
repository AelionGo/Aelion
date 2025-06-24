// Filename: ServiceContext.go
// Description: 服务上下文

// Package svc 服务上下文
package svc

import (
	"github.com/AelionGo/Aelion/internal/config"
)

type ServiceContext struct {
	Config *config.Config
}

func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
