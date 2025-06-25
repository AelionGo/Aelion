package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

		db, err = gorm.Open(sqlite.Open(sqlFile), &gorm.Config{})
	}

	//MySQL
	if sqlType == "mysql" {
		host := os.Getenv("AL_DB_Host")
		port := os.Getenv("AL_DB_Port")
		name := os.Getenv("AL_DB_Name")
		username := os.Getenv("AL_DB_Username")
		password := os.Getenv("AL_DB_Password")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, name)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	//Postgres
	if sqlType == "postgres" {
		host := os.Getenv("AL_DB_Host")
		port := os.Getenv("AL_DB_Port")
		name := os.Getenv("AL_DB_Name")
		username := os.Getenv("AL_DB_Username")
		password := os.Getenv("AL_DB_Password")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, name)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&Config{})

	return err
}
