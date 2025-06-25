package config

// RegisterCaptchaEnabled 检查注册时是否启用图形验证码
func (c *Config) RegisterCaptchaEnabled() (bool, error) {
	res, err := get("register_captcha_enabled")
	if err != nil {
		return false, err
	}
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
		return false, err
	}
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
