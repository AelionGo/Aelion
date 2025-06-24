// Filename: config.go
// Description: 配置结构与初始化

// Package config 初始化和管理配置信息
package config

import "github.com/zeromicro/go-zero/core/conf"

type ServerConfig struct {
	Host string `json:",default=0.0.0.0"`
	Port int    `json:",default=8080"`
}

type AuthConfig struct {
	RegisterCaptchaEnabled bool
	LoginCaptchaEnabled    bool
}

type Config struct {
	Server ServerConfig
}

func Init(configFile string) (*Config, error) {
	var c Config
	var sc ServerConfig
	conf.MustLoad(configFile, &sc)
	c.Server = sc

	//TODO: 从MySQL加载其他配置

	return &c, nil
}
