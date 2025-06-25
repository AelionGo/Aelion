package models

import "sync"

const (
	PolicyTypeLocal = "local" // 本地策略
)

type Policy struct {
	Id   string `gorm:"column:id;type:varchar(64);primary_key;comment:策略ID"`
	Type string `gorm:"column:type;type:varchar(64);not null;comment:策略类型"`

	// TODO: 添加更多字段
}

type PolicyModel struct {
	mu sync.Mutex
}

func NewPolicyModel() *PolicyModel {
	return &PolicyModel{}
}

func (p *PolicyModel) TableName() string {
	return "policy"
}

func (p *PolicyModel) GetOneById(id int) (*Policy, error) {
	var policy Policy
	err := db.Where("id = ?", id).First(&policy).Error
	if err != nil {
		return nil, err
	}
	return &policy, nil
}

func (p *PolicyModel) Create(policy *Policy) error {
	err := db.Create(policy).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PolicyModel) UpdateById(id int, policy *Policy) error {
	err := db.Model(&Policy{}).Where("id = ?", id).Updates(policy).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PolicyModel) DeleteById(id int) error {
	err := db.Where("id = ?", id).Delete(&Policy{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PolicyModel) CountAll() (int64, error) {
	var count int64
	err := db.Model(&Policy{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *PolicyModel) GetFirst() (*Policy, error) {
	var policy Policy
	err := db.First(&policy).Error
	if err != nil {
		return nil, err
	}
	return &policy, nil
}
