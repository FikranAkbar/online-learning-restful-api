package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/core"
	"online-learning-restful-api/model/entity"
)

type MasterCategorySeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterCategorySeeder(cfg gormseeder.SeederConfiguration) *MasterCategorySeeder {
	return &MasterCategorySeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func (s *MasterCategorySeeder) Seed(db *gorm.DB) error {
	categories := []entity.MasterCategory{
		{
			EntityModel:  core.EntityModel{ID: 1},
			CategoryName: "Entrepreneurship",
		},
		{
			EntityModel:  core.EntityModel{ID: 2},
			CategoryName: "Marketing",
		},
		{
			EntityModel:  core.EntityModel{ID: 3},
			CategoryName: "English",
		},
		{
			EntityModel:  core.EntityModel{ID: 4},
			CategoryName: "Creative",
		},
	}

	return db.CreateInBatches(categories, len(categories)).Error
}

func (s *MasterCategorySeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterCategory{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
