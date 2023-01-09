package database

import (
	"gorm.io/gorm"
	"log"
	"online-learning-restful-api/helper"
	. "online-learning-restful-api/model/entity"
)

func GetEntities() []any {
	return []any{
		&MasterAccount{},
		&MasterCategory{},
		&MasterCity{},
		&MasterCourse{},
		&MasterCourseComingSoon{},
		&MasterCourseReview{},
		&MasterCourseStatus{},
		&MasterCourseSummary{},
		&MasterDay{},
		&MasterElearningModule{},
		&MasterExpert{},
		&MasterIndustryInsight{},
		&MasterInsightAttachment{},
		&MasterLikeableType{},
		&MasterPaymentStatus{},
		&MasterPromo{},
		&MasterProvince{},
		&MasterQuiz{},
		&MasterUser{},
		&MasterUserType{},
		&MasterVideo{},
		&MasterWebinarSession{},
		&TrxCourseCategory{},
		&TrxCourseQnaAnswer{},
		&TrxCourseQnaQuestion{},
		&TrxElearningModuleSequence{},
		&TrxExpertCategory{},
		&TrxForgotPassword{},
		&TrxIndustryInsightAttachment{},
		&TrxQuizUserAnswer{},
		&TrxUserCart{},
		&TrxUserCourse{},
		&TrxUserLike{},
		&TrxUserPaymentHistory{},
		&TrxUserVideoProgression{},
		&TrxUserWebinarAttendance{},
		&TrxWebinarSessionSequence{},
	}
}

func Migrate(db *gorm.DB) {
	log.Printf("Migrate Start")

	entities := GetEntities()
	migrator := db.Migrator()

	for _, entity := range helper.ReverseSlice(entities) {
		dropTableErr := migrator.DropTable(entity)
		helper.PanicIfError(dropTableErr)
	}

	for _, entity := range entities {
		migrateTableErr := migrator.AutoMigrate(entity)
		helper.PanicIfError(migrateTableErr)
	}

	log.Printf("Migrate: Success")
}
