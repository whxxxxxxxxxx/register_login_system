package dao

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm" // 新版 GORM 的导入路径
	"gorm.io/gorm/logger"
	"log"
	"os"
	"register_log/model"
	"time"
)

var DB *gorm.DB

func Database(connString string) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出到标准输出）
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	migration()
}

func migration() {
	//自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&model.User{})
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}
