package test

import (
	"online-learning-restful-api/app/database"
	"online-learning-restful-api/app/database/seeder"
	"testing"
)

func TestMigrateClearSeedSequentiallyDB(t *testing.T) {
	t.Run("Test Migrate DB", func(t *testing.T) {
		database.Migrate(database.NewDB())
	})
	t.Run("Test Clear DB", func(t *testing.T) {
		seeder.ClearDB(database.NewDB())
	})
	t.Run("Test Seed DB", func(t *testing.T) {
		seeder.SeedDB(database.NewDB())
	})
	t.Run("Test Clear And Seed DB", func(t *testing.T) {
		seeder.ClearDB(database.NewDB())
		seeder.SeedDB(database.NewDB())
	})
}

func TestMigrateClearSeedSequentiallyDBTest(t *testing.T) {
	t.Run("Test Migrate DB", func(t *testing.T) {
		database.Migrate(database.NewTestDB())
	})
	t.Run("Test Clear DB", func(t *testing.T) {
		seeder.ClearDB(database.NewTestDB())
	})
	t.Run("Test Seed DB", func(t *testing.T) {
		seeder.SeedDB(database.NewTestDB())
	})
	t.Run("Test Clear And Seed DB", func(t *testing.T) {
		seeder.ClearDB(database.NewTestDB())
		seeder.SeedDB(database.NewTestDB())
	})
}
