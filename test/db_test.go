package test

import (
	"online-learning-restful-api/app/database"
	"online-learning-restful-api/app/database/seeder"
	"testing"
)

func TestMigrateDB(t *testing.T) {
	database.Migrate(database.NewDB())
}

func TestSeedDB(t *testing.T) {
	seeder.SeedDB(database.NewDB())
}

func TestClearDB(t *testing.T) {
	seeder.ClearDB(database.NewDB())
}
