package models

import "sync"

const (
	GroupTypeAdmin  = iota // 管理员
	GroupTypeNormal        // 普通用户
)

type Group struct {
	Id          string `gorm:"column:id;type:varchar(64);primary_key;comment:用户组ID"`
	Name        string `gorm:"column:name;type:varchar(64);not null;unique;comment:用户组名称"`
	Description string `gorm:"column:description;type:varchar(256);comment:用户组描述"`
	Type        int    `gorm:"column:type;type:int;not null;comment:用户组类型，0-管理员，1-普通用户"`
	Policy      string `gorm:"column:policy;type:text;comment:用户组策略，逗号分隔的列表"`
}

type GroupModel struct {
	mu sync.Mutex
}

func NewGroupModel() *GroupModel {
	return &GroupModel{}
}

func (g *GroupModel) TableName() string {
	return "user_group"
}

func (g *GroupModel) GetOneByID(id string) (*Group, error) {
	var group Group
	err := db.Where("id = ?", id).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (g *GroupModel) GetOneByName(name string) (*Group, error) {
	var group Group
	err := db.Where("name = ?", name).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (g *GroupModel) Create(group *Group) error {
	err := db.Create(group).Error
	if err != nil {
		return err
	}
	return nil
}

func (g *GroupModel) UpdateById(id string, group *Group) error {
	err := db.Model(&Group{}).Where("id = ?", id).Updates(group).Error
	if err != nil {
		return err
	}
	return nil
}

func (g *GroupModel) DeleteById(id string) error {
	err := db.Where("id = ?", id).Delete(&Group{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (g *GroupModel) CountAll() (int64, error) {
	var count int64
	err := db.Model(&Group{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (g *GroupModel) GetOneByType(groupType int) (*Group, error) {
	var group Group
	err := db.Where("type = ?", groupType).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}
