package seeder

import (
	"fmt"
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type TrxQuizUserAnswerSeeder struct {
	gormseeder.SeederAbstract
}

func NewTrxQuizUserAnswerSeeder(cfg gormseeder.SeederConfiguration) *TrxQuizUserAnswerSeeder {
	return &TrxQuizUserAnswerSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitQuizUserAnswerData() []entity.TrxQuizUserAnswer {
	var quizUserAnswers []entity.TrxQuizUserAnswer
	var i uint
	for i = 0; i < 48; i++ {
		data := entity.TrxQuizUserAnswer{
			ID:         i + 1,
			QuizId:     i + 1,
			UserId:     (i % 4) + 1,
			QuizAnswer: fmt.Sprintf("Quiz Answer Number %v", i+1),
		}

		quizUserAnswers = append(quizUserAnswers, data)
	}
	return quizUserAnswers
}

func (s *TrxQuizUserAnswerSeeder) Seed(db *gorm.DB) error {
	data := InitQuizUserAnswerData()
	return db.CreateInBatches(data, len(data)).Error
}

func (s *TrxQuizUserAnswerSeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxQuizUserAnswer{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
