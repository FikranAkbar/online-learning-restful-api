package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type TrxWebinarSessionSequenceSeeder struct {
	gormseeder.SeederAbstract
}

func NewTrxWebinarSessionSequenceSeeder(cfg gormseeder.SeederConfiguration) *TrxWebinarSessionSequenceSeeder {
	return &TrxWebinarSessionSequenceSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitWebinarSessionSequenceData() []entity.TrxWebinarSessionSequence {
	var elearningModuleSequences []entity.TrxWebinarSessionSequence
	var i uint
	for i = 0; i < 48; i++ {
		data := entity.TrxWebinarSessionSequence{
			WebinarSessionId: i + 1,
			OrderNumber:      (i % 4) + 1,
		}

		elearningModuleSequences = append(elearningModuleSequences, data)
	}
	return elearningModuleSequences
}

func (s *TrxWebinarSessionSequenceSeeder) Seed(db *gorm.DB) error {
	data := InitWebinarSessionSequenceData()
	return db.CreateInBatches(data, len(data)).Error
}

func (s *TrxWebinarSessionSequenceSeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxWebinarSessionSequence{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
