package database

import (
	"online-learning-restful-api/app/database/seeder"
	"testing"
)

func TestMigrateClearSeedSequentiallyDB(t *testing.T) {
	t.Run("Test Migrate DB", func(t *testing.T) {
		Migrate(NewDB())
	})
	t.Run("Test Clear DB", func(t *testing.T) {
		seeder.ClearDB(NewDB())
	})
	t.Run("Test Seed DB", func(t *testing.T) {
		seeder.SeedDB(NewDB())
	})
	t.Run("Test Clear And Seed DB", func(t *testing.T) {
		seeder.ClearDB(NewDB())
		seeder.SeedDB(NewDB())
	})
}

func TestMigrateClearSeedSequentiallyDBTest(t *testing.T) {
	t.Run("Test Migrate DB", func(t *testing.T) {
		Migrate(NewTestDB())
	})
	t.Run("Test Clear DB", func(t *testing.T) {
		seeder.ClearDB(NewTestDB())
	})
	t.Run("Test Seed DB", func(t *testing.T) {
		seeder.SeedDB(NewTestDB())
	})
	t.Run("Test Clear And Seed DB", func(t *testing.T) {
		seeder.ClearDB(NewTestDB())
		seeder.SeedDB(NewTestDB())
	})
}
