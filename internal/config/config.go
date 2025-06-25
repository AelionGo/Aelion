// Filename: config.go
// Description: 配置结构与初始化

// Package config 配置项的读取和设置
package config

import (
	"fmt"
	"github.com/AelionGo/Aelion/models"
	"github.com/zeromicro/go-zero/core/conf"
)

var (
	m     *models.ConfigModel
	cache map[string]string
)

type ServerConfig struct {
	Host string `json:",default=0.0.0.0"`
	Port int    `json:",default=8080"`
}

type Config struct {
	Server ServerConfig
}

func Init(configFile string) (*Config, error) {
	var c Config
	var sc ServerConfig
	conf.MustLoad(configFile, &sc)
	c.Server = sc

	m = models.NewConfigModel()
	cache = make(map[string]string)

	return &c, nil
}

func get(key string) (string, error) {
	//检查本地缓存
	res, ok := cache[key]
	if ok {
		return res, nil
	}

	res, err := m.GetOne(key)
	fmt.Println(res, err)
	if err != nil {
		return "", err
	}

	cache[key] = res
	return res, nil
}

func set(key, value string) error {
	delete(cache, key) // 删除缓存中的旧值
	return m.SetOne(key, value)
}
