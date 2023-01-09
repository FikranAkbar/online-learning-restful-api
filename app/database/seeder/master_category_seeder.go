package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/model/entity"
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
			Model:        gorm.Model{ID: 1},
			CategoryName: "Entrepreneurship",
		},
		{
			Model:        gorm.Model{ID: 2},
			CategoryName: "Marketing",
		},
		{
			Model:        gorm.Model{ID: 3},
			CategoryName: "English",
		},
		{
			Model:        gorm.Model{ID: 4},
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
