package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	// 连接 MySQL 数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动迁移 User 结构体对应的表
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate table")
	}
	// 自动迁移 Video 结构体对应的表
	err = db.AutoMigrate(&Video{})
	if err != nil {
		panic("failed to migrate table")
	}
	// 修改数据表名字
	if err := db.Migrator().RenameTable("users", "user_table"); err != nil {
		panic("failed to rename table")
	}
	if err := db.Migrator().RenameTable("videos", "videos_table"); err != nil {
		panic("failed to rename table")
	}
}
