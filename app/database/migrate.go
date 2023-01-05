package database

import (
	"log"
)

func Migrate() {
	log.Printf("Migrate Start")
	db := NewDB()

	db.AutoMigrate()
}
