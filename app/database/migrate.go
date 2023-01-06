package database

import (
	"gorm.io/gorm"
	"log"
	"online-learning-restful-api/helper"
	. "online-learning-restful-api/model/entity"
)

func Migrate(db *gorm.DB) {
	log.Printf("Migrate Start")

	helper.PanicIfError(db.SetupJoinTable(&MasterExpert{}, "Categories", &TrxExpertCategory{}))

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
		&MasterUser{},
		&MasterUserType{},
		&MasterVideo{},
		&MasterWebinarSession{},
		//&TrxCourseCategory{},
		//&TrxCourseQnaAnswer{},
		//&TrxCourseQnaQuestion{},
		//&TrxElearningModuleSequence{},
		&TrxExpertCategory{},
		//&TrxForgotPassword{},
		//&TrxIndustryInsightAttachment{},
		//&TrxQuizUserAnswer{},
		//&TrxUserCart{},
		//&TrxUserCourse{},
		//&TrxUserLike{},
		//&TrxUserPaymentHistory{},
		//&TrxUserVideoProgression{},
		//&TrxUserWebinarAttendance{},
		//&TrxWebinarSessionSequence{},
	)
	helper.PanicIfError(err)

	log.Printf("Migrate: Success")
}
