package main

import (
	"feed/internal/database/api"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// db := api.InitialDataBase() // initialize database 只需要一次
	dsn := "root:root@tcp(101.33.249.244:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	api.TestDateBase(db)
}
