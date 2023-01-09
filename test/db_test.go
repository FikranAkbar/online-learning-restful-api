package test

import (
	"online-learning-restful-api/app/database"
	"online-learning-restful-api/app/database/seeder"
	"testing"
)

func TestDBMigrateClearSeedSequentially(t *testing.T) {
	t.Run("Test Migrate DB", func(t *testing.T) {
		database.Migrate(database.NewDB())
	})
	t.Run("Test Clear DB", func(t *testing.T) {
		seeder.ClearDB(database.NewDB())
	})
	t.Run("Test Seed DB", func(t *testing.T) {
		seeder.SeedDB(database.NewDB())
	})
}
