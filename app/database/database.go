package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"online-learning-restful-api/config"
	"online-learning-restful-api/helper"
)

func NewDB(config *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.CreateDBURL()), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	helper.PanicIfError(err)

	return db
}

func NewTestDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/online_learning_database_test?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	helper.PanicIfError(err)

	return db
}
