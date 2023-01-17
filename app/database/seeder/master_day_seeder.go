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
			ID:   1,
			Name: "Monday",
		},
		{
			ID:   2,
			Name: "Tuesday",
		},
		{
			ID:   3,
			Name: "Wednesday",
		},
		{
			ID:   4,
			Name: "Thursday",
		},
		{
			ID:   5,
			Name: "Friday",
		},
		{
			ID:   6,
			Name: "Saturday",
		},
		{
			ID:   7,
			Name: "Sunday",
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
