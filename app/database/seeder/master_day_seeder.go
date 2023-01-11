package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterDaySeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterDaySeeder(cfg gormseeder.SeederConfiguration) *MasterDaySeeder {
	return &MasterDaySeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitDayData() []entity.MasterDay {
	return []entity.MasterDay{
		{
			Model: gorm.Model{ID: 1},
			Name:  "Monday",
		},
		{
			Model: gorm.Model{ID: 2},
			Name:  "Tuesday",
		},
		{
			Model: gorm.Model{ID: 3},
			Name:  "Wednesday",
		},
		{
			Model: gorm.Model{ID: 4},
			Name:  "Thursday",
		},
		{
			Model: gorm.Model{ID: 5},
			Name:  "Friday",
		},
		{
			Model: gorm.Model{ID: 6},
			Name:  "Saturday",
		},
		{
			Model: gorm.Model{ID: 7},
			Name:  "Sunday",
		},
	}
}

func (s *MasterDaySeeder) Seed(db *gorm.DB) error {
	data := InitDayData()
	return db.CreateInBatches(data, len(data)).Error
}

func (s *MasterDaySeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterDay{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
