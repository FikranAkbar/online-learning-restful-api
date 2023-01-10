package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"time"
)

type TrxExpertCategorySeeder struct {
	gormseeder.SeederAbstract
}

func NewTrxExpertCategorySeeder(cfg gormseeder.SeederConfiguration) *TrxExpertCategorySeeder {
	return &TrxExpertCategorySeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitExpertCategoryData() []entity.TrxExpertCategory {
	return []entity.TrxExpertCategory{
		{
			ExpertId:   1,
			CategoryId: 1,
			CreatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
		},
		{
			ExpertId:   1,
			CategoryId: 2,
			CreatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
		},
		{
			ExpertId:   2,
			CategoryId: 2,
			CreatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
		},
		{
			ExpertId:   2,
			CategoryId: 4,
			CreatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
		},
		{
			ExpertId:   3,
			CategoryId: 3,
			CreatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
		},
		{
			ExpertId:   3,
			CategoryId: 1,
			CreatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
		},
		{
			ExpertId:   4,
			CategoryId: 4,
			CreatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
		},
		{
			ExpertId:   4,
			CategoryId: 3,
			CreatedAt:  time.Time{},
			DeletedAt:  gorm.DeletedAt{},
		},
	}
}

func (s *TrxExpertCategorySeeder) Seed(db *gorm.DB) error {
	expertCategories := InitExpertCategoryData()

	return db.CreateInBatches(expertCategories, len(expertCategories)).Error
}

func (s *TrxExpertCategorySeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxExpertCategory{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
