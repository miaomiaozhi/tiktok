package manager

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Manager struct {
	handle *gorm.DB
}

func (s *Manager) NewManager(db *gorm.DB) {
	s.handle = db
}

var (
	M Manager
)

const (
	user         = "root"
	password     = "root"
	host         = "127.0.0.1"
	port         = "3306"
	databaseName = "tiktok"
)

func init() {
	// database 在本地，因此使用本地ip
	// dsn := "root:root@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             200,         // 慢 SQL 阈值
			LogLevel:                  logger.Warn, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	db, err := gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				user,
				password,
				host,
				port,
				databaseName,
			),
		),
		&gorm.Config{
			Logger: newLogger,
		},
	)
	if err != nil {
		log.Println(err)
	}
	M.NewManager(db)
}
