package internal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"ocean_mall/account/account_srv/model"
	"os"
	"time"
)

var DB *gorm.DB
var err error

func InitDB() {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		LogLevel:                  logger.Info,
		SlowThreshold:             time.Second,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "ocean_mall")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	err = DB.AutoMigrate(&model.Account{})
	if err != nil {
		fmt.Println(err)
	}

}
