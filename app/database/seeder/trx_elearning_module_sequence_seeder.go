package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type TrxElearningModuleSequenceSeeder struct {
	gormseeder.SeederAbstract
}

func NewTrxElearningModuleSequenceSeeder(cfg gormseeder.SeederConfiguration) *TrxElearningModuleSequenceSeeder {
	return &TrxElearningModuleSequenceSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitElearningModuleSequenceData() []entity.TrxElearningModuleSequence {
	var elearningModuleSequences []entity.TrxElearningModuleSequence
	var i uint
	for i = 0; i < 48; i++ {
		data := entity.TrxElearningModuleSequence{
			ModuleId:    i + 1,
			OrderNumber: (i % 4) + 1,
		}

		elearningModuleSequences = append(elearningModuleSequences, data)
	}
	return elearningModuleSequences
}

func (s *TrxElearningModuleSequenceSeeder) Seed(db *gorm.DB) error {
	data := InitElearningModuleSequenceData()
	return db.CreateInBatches(data, len(data)).Error
}

func (s *TrxElearningModuleSequenceSeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxElearningModuleSequence{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
