package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type TrxCourseQnaAnswerSeeder struct {
	gormseeder.SeederAbstract
}

func NewTrxCourseQnaAnswerSeeder(cfg gormseeder.SeederConfiguration) *TrxCourseQnaAnswerSeeder {
	return &TrxCourseQnaAnswerSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitCourseQnaAnswerData() entity.TrxCourseQnaAnswer {
	return entity.TrxCourseQnaAnswer{
		ID:            uint(1),
		QnaQuestionId: uint(1),
		UserId:        uint(1),
		UserTypeId:    uint(1),
		Answer:        "Ini jawaban saya!",
	}
}

func (s *TrxCourseQnaAnswerSeeder) Seed(db *gorm.DB) error {
	data := InitCourseQnaAnswerData()

	return db.Create(&data).Error
}

func (s *TrxCourseQnaAnswerSeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxCourseQnaAnswer{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
