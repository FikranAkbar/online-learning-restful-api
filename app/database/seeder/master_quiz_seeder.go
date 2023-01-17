package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterQuizSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterQuizSeeder(cfg gormseeder.SeederConfiguration) *MasterQuizSeeder {
	return &MasterQuizSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitQuizData() []entity.MasterQuiz {
	var quizData []entity.MasterQuiz

	for i := 1; i <= 48; i++ {
		data := entity.MasterQuiz{
			ID:           uint(i),
			ModuleId:     uint(i),
			QuizQuestion: "Ini adalah deskripsi dari pertanyaan quiz. Berikut adalah pertanyaannya dan silahkan tuliskan jawaban quiznya, bagaimana pendapat anda mengenai materi yang baru saja anda pelajari ?",
			IsPublished:  false,
		}

		quizData = append(quizData, data)
	}

	return quizData
}

func (s *MasterQuizSeeder) Seed(db *gorm.DB) error {
	quizData := InitQuizData()

	return db.CreateInBatches(quizData, len(quizData)).Error
}

func (s *MasterQuizSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterQuiz{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
