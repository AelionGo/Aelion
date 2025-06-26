package models

import (
	"fmt"
	"github.com/AelionGo/Aelion/pkg/hash"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var db *gorm.DB

func InitDB() error {
	//获取数据库类型
	sqlType := os.Getenv("AL_DB_Type")

	var err error

	//默认使用SQLite
	if sqlType == "" {
		//获取用户指定的文件路径
		sqlFile := os.Getenv("AL_DB_File")
		if sqlFile == "" {
			sqlFile = "al.db" //默认文件名
		}

		db, err = gorm.Open(sqlite.Open(sqlFile), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), // 禁用日志输出
		})
	}

	//MySQL
	if sqlType == "mysql" {
		host := os.Getenv("AL_DB_Host")
		port := os.Getenv("AL_DB_Port")
		name := os.Getenv("AL_DB_Name")
		username := os.Getenv("AL_DB_Username")
		password := os.Getenv("AL_DB_Password")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, name)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), // 禁用日志输出
		})
	}

	//Postgres
	if sqlType == "postgres" {
		host := os.Getenv("AL_DB_Host")
		port := os.Getenv("AL_DB_Port")
		name := os.Getenv("AL_DB_Name")
		username := os.Getenv("AL_DB_Username")
		password := os.Getenv("AL_DB_Password")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, name)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), // 禁用日志输出
		})
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&Config{}, &User{}, &Group{}, &Policy{})

	var policyId, adminGroupId, normalGroupId string

	p := NewPolicyModel()
	count, err := p.CountAll()
	if err != nil {
		return fmt.Errorf("failed to count policies: %w", err)
	}
	if count == 0 {
		// 如果没有策略数据，则插入默认策略
		policyId = uuid.New().String()
		defaultPolicy := &Policy{
			Id:   policyId,
			Type: PolicyTypeLocal,
		}
		err = p.Create(defaultPolicy)
		if err != nil {
			return fmt.Errorf("failed to create default policy: %w", err)
		}
	} else {
		// 获取第一个策略的ID
		policy, err := p.GetFirst()
		if err != nil {
			return fmt.Errorf("failed to get default policy: %w", err)
		}
		policyId = policy.Id
	}

	g := NewGroupModel()
	count, err = g.CountAll()
	if err != nil {
		return fmt.Errorf("failed to count groups: %w", err)
	}
	if count == 0 {
		//创建初始用户组
		adminGroupId = uuid.New().String()
		group := &Group{
			Id:          adminGroupId,
			Name:        "default admin group",
			Description: "默认管理员组",
			Type:        GroupTypeAdmin, // 默认是管理员组
			Policy:      policyId,       // 关联默认策略
		}
		err = g.Create(group)
		if err != nil {
			return fmt.Errorf("failed to create default group: %w", err)
		}

		normalGroupId = uuid.New().String()
		group = &Group{
			Id:          normalGroupId,
			Name:        "default normal group",
			Description: "默认普通用户组",
			Type:        GroupTypeNormal, // 默认是普通用户组
			Policy:      policyId,        // 关联默认策略
		}
		err = g.Create(group)
		if err != nil {
			return fmt.Errorf("failed to create default normal group: %w", err)
		}
	} else {
		// 获取默认管理员组
		group, err := g.GetOneByType(GroupTypeAdmin)
		if err != nil {
			return fmt.Errorf("failed to get default admin group: %w", err)
		}
		adminGroupId = group.Id
	}

	u := NewUserModel()
	count, err = u.CountAll()
	if err != nil {
		return fmt.Errorf("failed to count users: %w", err)
	}
	if count == 0 {
		// 创建初始管理员用户
		userId := uuid.New().String()
		password := hash.GenerateRandomPassword()
		fmt.Println("初始管理员邮箱: admin@aelion.org")
		fmt.Println("初始管理员密码:", password)
		hashed, err := hash.HashPassword(password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		adminUser := &User{
			Id:       userId,
			Email:    "admin@aelion.org", // 初始管理员邮箱
			Phone:    "",                 // 初始管理员手机号
			Password: hashed,             // 初始管理员密码
			Nickname: "Admin",            // 初始管理员昵称
			Avatar:   "",                 // 初始管理员头像
			Status:   UserActive,         // 初始管理员状态正常
			Group:    adminGroupId,       // 关联默认管理员组
		}
		err = u.Create(adminUser)
		if err != nil {
			return fmt.Errorf("failed to create default admin user: %w", err)
		}
	}

	return err
}
