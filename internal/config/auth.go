// Filename: auth.go
// Description: 认证相关配置项的读取和设置

// Package config 配置项的读取和设置
package config

// RegisterCaptchaEnabled 检查注册时是否启用图形验证码
func (c *Config) RegisterCaptchaEnabled() (bool, error) {
	res, err := get("register_captcha_enabled")
	if err != nil {
		return false, set("register_captcha_enabled", "false")
	} // 如果获取失败，默认设置为 false
	if res == "true" {
		return true, nil
	}
	return false, nil
}

// SetRegisterCaptchaEnabled 设置注册时是否启用图形验证码
func (c *Config) SetRegisterCaptchaEnabled(enabled bool) error {
	if enabled {
		return set("register_captcha_enabled", "true")
	}
	return set("register_captcha_enabled", "false")
}

// LoginCaptchaEnabled 检查登录时是否启用图形验证码
func (c *Config) LoginCaptchaEnabled() (bool, error) {
	res, err := get("login_captcha_enabled")
	if err != nil {
		return false, set("login_captcha_enabled", "false")
	} // 如果获取失败，默认设置为 false
	if res == "true" {
		return true, nil
	}
	return false, nil
}

// SetLoginCaptchaEnabled 设置登录时是否启用图形验证码
func (c *Config) SetLoginCaptchaEnabled(enabled bool) error {
	if enabled {
		return set("login_captcha_enabled", "true")
	}
	return set("login_captcha_enabled", "false")
}
