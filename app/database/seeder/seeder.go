package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"log"
	"online-learning-restful-api/helper"
)

func getSeeders() []gormseeder.SeederInterface {
	return []gormseeder.SeederInterface{
		NewMasterCourseComingSoonSeeder(gormseeder.SeederConfiguration{}),
		NewMasterLikeableTypeSeeder(gormseeder.SeederConfiguration{}),
		NewMasterPaymentStatusSeeder(gormseeder.SeederConfiguration{}),
		NewMasterProvinceSeeder(gormseeder.SeederConfiguration{}),
		NewMasterCitySeeder(gormseeder.SeederConfiguration{}),
		NewMasterDaySeeder(gormseeder.SeederConfiguration{}),
		NewMasterUserTypeSeeder(gormseeder.SeederConfiguration{}),
		NewMasterAccountSeeder(gormseeder.SeederConfiguration{}),
		NewMasterUserSeeder(gormseeder.SeederConfiguration{}),
		NewMasterExpertSeeder(gormseeder.SeederConfiguration{}),
		NewMasterCategorySeeder(gormseeder.SeederConfiguration{}),
		NewMasterCourseStatusSeeder(gormseeder.SeederConfiguration{}),
		NewMasterCourseSeeder(gormseeder.SeederConfiguration{}),
		NewMasterCourseSummarySeeder(gormseeder.SeederConfiguration{}),
		NewMasterIndustryInsightSeeder(gormseeder.SeederConfiguration{}),
		NewMasterInsightAttachmentSeeder(gormseeder.SeederConfiguration{}),
		NewMasterElearningModuleSeeder(gormseeder.SeederConfiguration{}),
		NewMasterQuizSeeder(gormseeder.SeederConfiguration{}),
		NewMasterWebinarSessionSeeder(gormseeder.SeederConfiguration{}),
		NewTrxExpertCategorySeeder(gormseeder.SeederConfiguration{}),
		NewTrxCourseCategorySeeder(gormseeder.SeederConfiguration{}),
		NewTrxUserCourseSeeder(gormseeder.SeederConfiguration{}),
		NewTrxCourseReviewSeeder(gormseeder.SeederConfiguration{}),
		NewMasterVideoSeeder(gormseeder.SeederConfiguration{}),
		NewTrxIndustryInsightAttachmentSeeder(gormseeder.SeederConfiguration{}),
		NewTrxElearningModuleSequenceSeeder(gormseeder.SeederConfiguration{}),
		NewTrxWebinarSessionSequenceSeeder(gormseeder.SeederConfiguration{}),
		NewTrxQuizUserAnswerSeeder(gormseeder.SeederConfiguration{}),
		NewTrxUserLikeSeeder(gormseeder.SeederConfiguration{}),
		NewTrxCourseQnaQuestionSeeder(gormseeder.SeederConfiguration{}),
		NewTrxCourseQnaAnswerSeeder(gormseeder.SeederConfiguration{}),
		NewMasterPaymentMethodSeeder(gormseeder.SeederConfiguration{}),
	}

}

func SeedDB(db *gorm.DB) {
	log.Printf("Seed: Start")

	seedersStack := gormseeder.NewSeedersStack(db)
	for _, modelSeeder := range getSeeders() {
		seedersStack.AddSeeder(modelSeeder)
	}
	err := seedersStack.Seed()
	helper.PanicIfError(err)

	log.Printf("Seed: Success")
}

func ClearDB(db *gorm.DB) {
	log.Printf("Clear: Start")

	seedersStack := gormseeder.NewSeedersStack(db)
	for _, modelSeeder := range helper.ReverseSeeders(getSeeders()) {
		seedersStack.AddSeeder(modelSeeder)
	}
	err := seedersStack.Clear()
	helper.PanicIfError(err)

	log.Printf("Clear: Success")
}
