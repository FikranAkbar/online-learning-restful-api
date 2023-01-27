package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"online-learning-restful-api/config"
	"online-learning-restful-api/helper"
)

func NewDB() *gorm.DB {
	dbConfig := config.BuildDBConfig()
	db, err := gorm.Open(mysql.Open(dbConfig.DBUrl()), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	helper.PanicIfError(err)

	return db
}

func NewTestDB() *gorm.DB {
	dbConfig := config.BuildDBConfigForTest()
	db, err := gorm.Open(mysql.Open(dbConfig.DBUrl()), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	helper.PanicIfError(err)

	return db
}
