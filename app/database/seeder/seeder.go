package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"log"
	"online-learning-restful-api/helper"
)

func SeedDB(db *gorm.DB) {
	log.Printf("Seed: Start")

	seedersStack := gormseeder.NewSeedersStack(db)

	seeders := []gormseeder.SeederInterface{
		NewMasterUserTypeSeeder(gormseeder.SeederConfiguration{}),
		NewMasterAccountSeeder(gormseeder.SeederConfiguration{}),
		NewMasterExpertSeeder(gormseeder.SeederConfiguration{}),
		NewMasterCategorySeeder(gormseeder.SeederConfiguration{}),
		NewTrxExpertCategorySeeder(gormseeder.SeederConfiguration{}),
		NewMasterCourseStatusSeeder(gormseeder.SeederConfiguration{}),
		NewMasterCourseSeeder(gormseeder.SeederConfiguration{}),
		NewTrxCourseCategorySeeder(gormseeder.SeederConfiguration{}),
	}

	for _, modelSeeder := range seeders {
		seedersStack.AddSeeder(modelSeeder)
	}

	seedErr := seedersStack.Seed()
	helper.PanicIfError(seedErr)

	log.Printf("Seed: Success")
}

func ClearDB(db *gorm.DB) {
	log.Printf("Clear: Start")

	seedersStack := gormseeder.NewSeedersStack(db)

	seeders := []gormseeder.SeederInterface{
		NewMasterUserTypeSeeder(gormseeder.SeederConfiguration{}),
		NewMasterAccountSeeder(gormseeder.SeederConfiguration{}),
		NewMasterExpertSeeder(gormseeder.SeederConfiguration{}),
		NewMasterCategorySeeder(gormseeder.SeederConfiguration{}),
		NewTrxExpertCategorySeeder(gormseeder.SeederConfiguration{}),
	}

	for _, modelSeeder := range seeders {
		seedersStack.AddSeeder(modelSeeder)
	}

	seedErr := seedersStack.Clear()
	helper.PanicIfError(seedErr)

	log.Printf("Clear: Success")
}
