package database

import (
	"online-learning-restful-api/app/database/seeder"
	"online-learning-restful-api/config"
	"testing"
)

func TestMigrateClearSeedSequentiallyDB(t *testing.T) {
	c := &config.Config{
		DBHost:     "localhost",
		DBPort:     3306,
		DBUser:     "root",
		DBName:     "online_learning_database",
		DBPassword: "password",
	}
	t.Run("Test Migrate DB", func(t *testing.T) {
		Migrate(NewDB(c))
	})
	t.Run("Test Clear DB", func(t *testing.T) {
		seeder.ClearDB(NewDB(c))
	})
	t.Run("Test Seed DB", func(t *testing.T) {
		seeder.SeedDB(NewDB(c))
	})
	t.Run("Test Clear And Seed DB", func(t *testing.T) {
		seeder.ClearDB(NewDB(c))
		seeder.SeedDB(NewDB(c))
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
