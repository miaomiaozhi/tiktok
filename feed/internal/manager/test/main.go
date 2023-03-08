package main

import (
	manager "feed/internal/manager"
)

func main() {
	// db := api.InitialDataBase() // initialize database 只需要一次
	// database 在本地，因此使用本地ip，
	// dsn := "root:root@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Println(err)
	// }
	// api.TestDateBase(db)

	// 删除表
	// db.Exec("drop table if exists users;")
	// db.Exec("drop table if exists videos;")

	// // 创建映射关系，并且将 feed_user -> users
	// db.Table("users").AutoMigrate(&api.User{})
	// db.Table("videos").AutoMigrate(&api.Video{})

	// // 插入
	// user := api.User {
	// 	Id:   123,
	// 	Name: "test user",
	// }
	// if err := db.Table("users").Create(&user).Error; err != nil {
	// 	log.Fatalf("Failed to insert a user: %v", err)
	// }

	// 查询
	manager.M.TestQueryUser(123)
}
