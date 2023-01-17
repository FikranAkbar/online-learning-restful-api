package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterCategorySeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterCategorySeeder(cfg gormseeder.SeederConfiguration) *MasterCategorySeeder {
	return &MasterCategorySeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitCategoryData() []entity.MasterCategory {
	return []entity.MasterCategory{
		{
			ID:           1,
			CategoryName: "Entrepreneurship",
		},
		{
			ID:           2,
			CategoryName: "Marketing",
		},
		{
			ID:           3,
			CategoryName: "English",
		},
		{
			ID:           4,
			CategoryName: "Creative",
		},
	}
}

func (s *MasterCategorySeeder) Seed(db *gorm.DB) error {
	categories := InitCategoryData()

	return db.CreateInBatches(categories, len(categories)).Error
}

func (s *MasterCategorySeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterCategory{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
