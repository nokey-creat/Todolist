package database

import (
	"Todolist/config"
	"Todolist/global"
	"fmt"
	"log"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.Database.User, config.AppConfig.Database.Password, config.AppConfig.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true}) //启用这个选项，将sql方言转为gorm error
	if err != nil {
		log.Fatalf("Init DB error: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Init DB error: %v", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	sqlDB.SetMaxIdleConns(config.AppConfig.Database.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.AppConfig.Database.MaxOpenConns)
	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(config.AppConfig.Database.ConnMaxLifetime))

	global.DB = db

}
