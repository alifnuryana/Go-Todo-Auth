package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alifnuryana/go-todo-auth/config"
	"github.com/alifnuryana/go-todo-auth/helper"
	"github.com/alifnuryana/go-todo-auth/model"
	mysqli "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
	mysqlConfig := mysqli.Config{
		User:                 config.Load("DB_USER"),
		Passwd:               config.Load("DB_PASS"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", config.Load("DB_HOST"), config.Load("DB_PORT")),
		DBName:               config.Load("DB_NAME"),
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})

	db, err := gorm.Open(mysql.Open(mysqlConfig.FormatDSN()), &gorm.Config{
		Logger: newLogger,
	})
	helper.FatalIfError(err)

	db.AutoMigrate(&model.Todo{})

	sqlDB, err := db.DB()
	helper.FatalIfError(err)

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
}
