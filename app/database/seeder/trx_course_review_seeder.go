package seeder

import (
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type TrxCourseReviewSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewTrxCourseReviewSeeder(cfg gorm_seeder.SeederConfiguration) *TrxCourseReviewSeeder {
	return &TrxCourseReviewSeeder{gorm_seeder.SeederAbstract{Configuration: cfg}}
}

func InitCourseReviewData() []entity.TrxCourseReview {
	return []entity.TrxCourseReview{
		//#Review for Intro to Enterpreneurship
		{
			Model:      gorm.Model{ID: 1},
			CourseId:   1,
			UserId:     1,
			ReviewDesc: "Coursenya mantaps, kesukaan gue banget",
			ReviewRate: 4.0,
		},
		{
			Model:      gorm.Model{ID: 2},
			CourseId:   1,
			UserId:     4,
			ReviewDesc: "Ajibb ajibbbb coursenya",
			ReviewRate: 4.7,
		},
		{
			Model:      gorm.Model{ID: 3},
			CourseId:   1,
			UserId:     3,
			ReviewDesc: "Penjelasan expertnya sangat mudah dimengerti",
			ReviewRate: 5.0,
		},
		//#

		//#Review for Intermediate to Enterpreneurship
		{
			Model:      gorm.Model{ID: 4},
			CourseId:   5,
			UserId:     3,
			ReviewDesc: "Modulnya sangat detail dan videonya mantap",
			ReviewRate: 4.2,
		},
		{
			Model:      gorm.Model{ID: 5},
			CourseId:   5,
			UserId:     2,
			ReviewDesc: "Expertnya favorit bangets",
			ReviewRate: 4.8,
		},
		{
			Model:      gorm.Model{ID: 6},
			CourseId:   5,
			UserId:     4,
			ReviewDesc: "Penjelasan materinya sangat detail dan quiznya bagus sekali",
			ReviewRate: 4.5,
		},
		//#

		//#Review for Expert to Enterpreneurship
		{
			Model:      gorm.Model{ID: 7},
			CourseId:   9,
			UserId:     1,
			ReviewDesc: "Sesi webinarnya sangat seru, expertnya asik banget",
			ReviewRate: 4.2,
		},
		{
			Model:      gorm.Model{ID: 8},
			CourseId:   9,
			UserId:     3,
			ReviewDesc: "Materinya sangat dalam sekali. Ini nih yang saya cari",
			ReviewRate: 4.8,
		},
		{
			Model:      gorm.Model{ID: 9},
			CourseId:   9,
			UserId:     2,
			ReviewDesc: "Daging semua asli dah ini course",
			ReviewRate: 4.5,
		},
		//#

		//#Review for Intro to Marketing
		{
			Model:      gorm.Model{ID: 10},
			CourseId:   2,
			UserId:     1,
			ReviewDesc: "Coursenya mantaps, kesukaan gue banget",
			ReviewRate: 4.0,
		},
		{
			Model:      gorm.Model{ID: 11},
			CourseId:   2,
			UserId:     4,
			ReviewDesc: "Ajibb ajibbbb coursenya",
			ReviewRate: 4.7,
		},
		{
			Model:      gorm.Model{ID: 12},
			CourseId:   2,
			UserId:     3,
			ReviewDesc: "Penjelasan expertnya sangat mudah dimengerti",
			ReviewRate: 5.0,
		},
		//#

		//#Review for Intermediate to Marketing
		{
			Model:      gorm.Model{ID: 13},
			CourseId:   6,
			UserId:     3,
			ReviewDesc: "Modulnya sangat detail dan videonya mantap",
			ReviewRate: 4.2,
		},
		{
			Model:      gorm.Model{ID: 14},
			CourseId:   6,
			UserId:     2,
			ReviewDesc: "Expertnya favorit bangets",
			ReviewRate: 4.8,
		},
		{
			Model:      gorm.Model{ID: 15},
			CourseId:   6,
			UserId:     4,
			ReviewDesc: "Penjelasan materinya sangat detail dan quiznya bagus sekali",
			ReviewRate: 4.5,
		},
		//#

		//#Review for Expert to Marketing
		{
			Model:      gorm.Model{ID: 16},
			CourseId:   10,
			UserId:     1,
			ReviewDesc: "Sesi webinarnya sangat seru, expertnya asik banget",
			ReviewRate: 4.2,
		},
		{
			Model:      gorm.Model{ID: 17},
			CourseId:   10,
			UserId:     3,
			ReviewDesc: "Materinya sangat dalam sekali. Ini nih yang saya cari",
			ReviewRate: 4.8,
		},
		{
			Model:      gorm.Model{ID: 18},
			CourseId:   10,
			UserId:     2,
			ReviewDesc: "Daging semua asli dah ini course",
			ReviewRate: 4.5,
		},
		//#

		//#Review for Intro to English
		{
			Model:      gorm.Model{ID: 19},
			CourseId:   3,
			UserId:     1,
			ReviewDesc: "Coursenya mantaps, kesukaan gue banget",
			ReviewRate: 4.0,
		},
		{
			Model:      gorm.Model{ID: 20},
			CourseId:   3,
			UserId:     4,
			ReviewDesc: "Ajibb ajibbbb coursenya",
			ReviewRate: 4.7,
		},
		{
			Model:      gorm.Model{ID: 21},
			CourseId:   3,
			UserId:     3,
			ReviewDesc: "Penjelasan expertnya sangat mudah dimengerti",
			ReviewRate: 5.0,
		},
		//#

		//#Review for Intermediate to English
		{
			Model:      gorm.Model{ID: 22},
			CourseId:   7,
			UserId:     3,
			ReviewDesc: "Modulnya sangat detail dan videonya mantap",
			ReviewRate: 4.2,
		},
		{
			Model:      gorm.Model{ID: 23},
			CourseId:   7,
			UserId:     2,
			ReviewDesc: "Expertnya favorit bangets",
			ReviewRate: 4.8,
		},
		{
			Model:      gorm.Model{ID: 24},
			CourseId:   7,
			UserId:     4,
			ReviewDesc: "Penjelasan materinya sangat detail dan quiznya bagus sekali",
			ReviewRate: 4.5,
		},
		//#

		//#Review for Expert to English
		{
			Model:      gorm.Model{ID: 25},
			CourseId:   11,
			UserId:     1,
			ReviewDesc: "Sesi webinarnya sangat seru, expertnya asik banget",
			ReviewRate: 4.2,
		},
		{
			Model:      gorm.Model{ID: 26},
			CourseId:   11,
			UserId:     3,
			ReviewDesc: "Materinya sangat dalam sekali. Ini nih yang saya cari",
			ReviewRate: 4.8,
		},
		{
			Model:      gorm.Model{ID: 27},
			CourseId:   11,
			UserId:     2,
			ReviewDesc: "Daging semua asli dah ini course",
			ReviewRate: 4.5,
		},
		//#

		//#Review for Intro to Creative
		{
			Model:      gorm.Model{ID: 28},
			CourseId:   12,
			UserId:     1,
			ReviewDesc: "Coursenya mantaps, kesukaan gue banget",
			ReviewRate: 4.0,
		},
		{
			Model:      gorm.Model{ID: 29},
			CourseId:   12,
			UserId:     4,
			ReviewDesc: "Ajibb ajibbbb coursenya",
			ReviewRate: 4.7,
		},
		{
			Model:      gorm.Model{ID: 30},
			CourseId:   12,
			UserId:     3,
			ReviewDesc: "Penjelasan expertnya sangat mudah dimengerti",
			ReviewRate: 5.0,
		},
		//#

		//#Review for Intermediate to Creative
		{
			Model:      gorm.Model{ID: 31},
			CourseId:   12,
			UserId:     3,
			ReviewDesc: "Modulnya sangat detail dan videonya mantap",
			ReviewRate: 4.2,
		},
		{
			Model:      gorm.Model{ID: 32},
			CourseId:   12,
			UserId:     2,
			ReviewDesc: "Expertnya favorit bangets",
			ReviewRate: 4.8,
		},
		{
			Model:      gorm.Model{ID: 33},
			CourseId:   12,
			UserId:     4,
			ReviewDesc: "Penjelasan materinya sangat detail dan quiznya bagus sekali",
			ReviewRate: 4.5,
		},
		//#

		//#Review for Expert to Creative
		{
			Model:      gorm.Model{ID: 34},
			CourseId:   12,
			UserId:     1,
			ReviewDesc: "Sesi webinarnya sangat seru, expertnya asik banget",
			ReviewRate: 4.2,
		},
		{
			Model:      gorm.Model{ID: 35},
			CourseId:   12,
			UserId:     3,
			ReviewDesc: "Materinya sangat dalam sekali. Ini nih yang saya cari",
			ReviewRate: 4.8,
		},
		{
			Model:      gorm.Model{ID: 36},
			CourseId:   12,
			UserId:     2,
			ReviewDesc: "Daging semua asli dah ini course",
			ReviewRate: 4.5,
		},
		//#
	}
}

func (s *TrxCourseReviewSeeder) Seed(db *gorm.DB) error {
	courseReviews := InitCourseReviewData()

	return db.CreateInBatches(courseReviews, len(courseReviews)).Error
}

func (s *TrxCourseReviewSeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxCourseReview{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
