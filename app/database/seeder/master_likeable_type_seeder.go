package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterLikeableTypeSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterLikeableTypeSeeder(cfg gormseeder.SeederConfiguration) *MasterLikeableTypeSeeder {
	return &MasterLikeableTypeSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitLikeableTypeData() []entity.MasterLikeableType {
	return []entity.MasterLikeableType{
		{
			Model:        gorm.Model{ID: 1},
			LikeableName: "quiz_answer",
		},
	}
}

func (s *MasterLikeableTypeSeeder) Seed(db *gorm.DB) error {
	data := InitLikeableTypeData()
	return db.CreateInBatches(data, len(data)).Error
}

func (s *MasterLikeableTypeSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterLikeableType{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
