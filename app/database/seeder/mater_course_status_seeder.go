package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/model/entity"
)

type MasterCourseStatusSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterCourseStatusSeeder(cfg gormseeder.SeederConfiguration) *MasterCourseStatusSeeder {
	return &MasterCourseStatusSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func (s *MasterCourseStatusSeeder) Seed(db *gorm.DB) error {
	courseStatuses := []entity.MasterCourseStatus{
		{
			Model: gorm.Model{ID: 1},
			Name:  "not_started",
		},
		{
			Model: gorm.Model{ID: 2},
			Name:  "ongoing",
		},
		{
			Model: gorm.Model{ID: 3},
			Name:  "completed",
		},
	}

	return db.CreateInBatches(courseStatuses, len(courseStatuses)).Error
}

func (s *MasterCourseStatusSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterCourseStatus{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
