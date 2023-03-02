package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
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

	// // 创建一个用户
	// user := User{Name: "Tom"}
	// result := db.Create(&user)
	// if result.Error != nil {
	// 	panic("failed to create user")
	// }

	// // 查询所有用户
	// var users []User
	// result = db.Find(&users)
	// if result.Error != nil {
	// 	panic("failed to query users")
	// }

	// // 输出查询结果
	// for _, u := range users {
	// 	fmt.Println(u.ID, u.Name)
	// }
}
