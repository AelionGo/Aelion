package config

import "github.com/AelionGo/Aelion/models"

// EmailValidationEnabled 是否启用邮箱激活
func (c *Config) EmailValidationEnabled() (bool, error) {
	res, err := get("email_validation_enabled")
	if err != nil {
		return false, set("email_validation_enabled", "false")
	} // 如果获取失败，默认设置为 false
	if res == "true" {
		return true, nil
	}
	return false, nil
}

// SetEmailValidationEnabled 设置邮箱激活启用状态
func (c *Config) SetEmailValidationEnabled(enabled bool) error {
	if enabled {
		return set("email_validation_enabled", "true")
	}
	return set("email_validation_enabled", "false")
}

// DefaultGroup 获取默认用户组
func (c *Config) DefaultGroup() (string, error) {
	res, err := get("default_group")
	if err != nil {
		g := models.NewGroupModel()
		group, err := g.GetOneByType(models.GroupTypeNormal) //获取一个普通用户组
		if err != nil {
			return "", err
		}
		return group.Id, set("default_group", group.Id)
	}
	return res, nil
}

// SetDefaultGroup 设置默认用户组
func (c *Config) SetDefaultGroup(group string) error {
	return set("default_group", group)
}
