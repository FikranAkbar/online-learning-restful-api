package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
)

func NewDB() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/online_learning_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	helper.PanicIfError(err)

	return db
}
