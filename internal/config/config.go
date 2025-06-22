// package config 初始化和管理配置信息

// Filename: config.go
// Description: 配置结构

package config

type ServerConfig struct {
	Host string `json:",default=0.0.0.0"`
	Port int    `json:",default=8080"`
}

type Config struct {
	Server ServerConfig
}
