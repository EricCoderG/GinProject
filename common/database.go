package common

import (
	"GinProject/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := "localhost"
	port := "3306"
	database := "gin"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate database,err:" + err.Error())
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
