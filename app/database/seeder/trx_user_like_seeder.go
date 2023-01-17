package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type TrxUserLikeSeeder struct {
	gormseeder.SeederAbstract
}

func NewTrxUserLikeSeeder(cfg gormseeder.SeederConfiguration) *TrxUserLikeSeeder {
	return &TrxUserLikeSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitUserLikeData() []entity.TrxUserLike {
	return []entity.TrxUserLike{
		{
			ID:             1,
			LikeableId:     1,
			LikeableTypeId: 1,
			UserId:         1,
			IsLike:         true,
		},
	}
}

func (s *TrxUserLikeSeeder) Seed(db *gorm.DB) error {
	data := InitUserLikeData()
	return db.CreateInBatches(data, len(data)).Error
}

func (s *TrxUserLikeSeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxUserLike{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
