package models

import "sync"

const (
	UserActive         = iota //用户状态：0-正常
	UserNeedActivation        //用户状态：1-需要激活
	UserBanned                //用户状态：2-被封禁
)

type User struct {
	Id       string `gorm:"column:id;type:varchar(128);primary_key;comment:用户ID"`
	Email    string `gorm:"column:email;type:varchar(128);unique;comment:用户邮箱"`
	Phone    string `gorm:"column:phone;type:varchar(16);unique;comment:用户手机号"` //邮箱和手机号至少有一个，且分别唯一
	Password string `gorm:"column:password;type:varchar(128);not null;comment:用户密码"`
	Nickname string `gorm:"column:nickname;type:varchar(64);not null;comment:用户昵称"`
	Avatar   string `gorm:"column:avatar;type:varchar(256);comment:用户头像"`
	Status   int    `gorm:"column:status;type:int;comment:用户状态，0-正常，1-需要激活，2-被封禁"`
	Group    string `gorm:"column:group;type:varchar(64);comment:用户组"`
}

type UserModel struct {
	mu sync.Mutex
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (u *UserModel) TableName() string {
	return "user"
}

func (u *UserModel) GetOneByEmail(email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserModel) GetOneByPhone(phone string) (*User, error) {
	var user User
	err := db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserModel) GetOneByID(id string) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserModel) Create(user *User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserModel) UpdateByID(id string, user *User) error {
	err := db.Model(&User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserModel) DeleteByID(id string) error {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserModel) CountAll() (int64, error) {
	var count int64
	err := db.Model(&User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
