package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type TrxCourseQnaQuestionSeeder struct {
	gormseeder.SeederAbstract
}

func NewTrxCourseQnaQuestionSeeder(cfg gormseeder.SeederConfiguration) *TrxCourseQnaQuestionSeeder {
	return &TrxCourseQnaQuestionSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitCourseQnaQuestionData() entity.TrxCourseQnaQuestion {
	return entity.TrxCourseQnaQuestion{
		ID:       uint(1),
		CourseId: uint(1),
		UserId:   uint(1),
		Question: "Ini pertanyaan saya (?)",
	}
}

func (s *TrxCourseQnaQuestionSeeder) Seed(db *gorm.DB) error {
	data := InitCourseQnaQuestionData()

	return db.Create(&data).Error
}

func (s *TrxCourseQnaQuestionSeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxCourseQnaQuestion{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
