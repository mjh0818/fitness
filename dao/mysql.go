package dao

import (
	"fitness/entity"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		viper.GetString("datasource.username"),
		viper.GetString("datasource.password"),
		viper.GetString("datasource.host"),
		viper.GetString("datasource.port"),
		viper.GetString("datasource.database"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
	}

	err = DB.AutoMigrate(&entity.User{}, &entity.Admin{}, &entity.Coach{}, &entity.Area{}, &entity.AreaBook{}, &entity.CoachBook{})
	if err != nil {
		log.Panic("failed to automigrate database")
	}
}

// GetDb 获取Db操作数据库

func GetDB() *gorm.DB {
	return DB
}
