package database

import (
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/seeders"
	"online-learning-restful-api/helper"
)

func SeedDB(db *gorm.DB) {
	seedersStack := gorm_seeder.NewSeedersStack(db)
	seedersStack.AddSeeder(
		seeders.NewMasterAccountSeeder(gorm_seeder.SeederConfiguration{Rows: 10}),
	)

	truncateErr := seedersStack.Clear()
	helper.PanicIfError(truncateErr)

	seedErr := seedersStack.Seed()
	helper.PanicIfError(seedErr)
}
