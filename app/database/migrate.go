package database

import (
	"log"
	"online-learning-restful-api/helper"
	. "online-learning-restful-api/model/domain"
)

func Migrate() {
	log.Printf("Migrate Start")
	db := NewDB()

	err := db.AutoMigrate(&MasterAccount{})
	helper.PanicIfError(err)

	log.Printf("Migrate: Success")
}
