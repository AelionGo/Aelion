package models

type Config struct {
	Key   string `gorm:"column:config_key;type:varchar(64);primary_key;comment:配置键"`
	Value string `gorm:"column:config_value;type:varchar(256);comment:配置值"`
}

type ConfigModel struct {
}

func NewConfigModel() *ConfigModel {
	return &ConfigModel{}
}

func (c *ConfigModel) TableName() string {
	return "config"
}

func (c *ConfigModel) GetOne(key string) (string, error) {
	var config Config
	err := db.Where("config_key = ?", key).First(&config).Error
	if err != nil {
		return "", err
	}
	return config.Value, nil
}

func (c *ConfigModel) SetOne(key, value string) error {
	config := Config{Key: key, Value: value}
	err := db.Save(&config).Error
	return err
}
