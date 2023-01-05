package database

import (
	"log"
	"online-learning-restful-api/helper"
	. "online-learning-restful-api/model/domain"
)

func Migrate() {
	log.Printf("Migrate Start")
	db := NewDB()

	err := db.AutoMigrate(
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
		&MasterPaymentChannel{},
		&MasterPaymentStatus{},
		&MasterPromo{},
		&MasterProvince{},
		&MasterQuiz{},
	)
	helper.PanicIfError(err)

	log.Printf("Migrate: Success")
}
