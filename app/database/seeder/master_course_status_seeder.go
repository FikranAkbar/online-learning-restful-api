package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterCourseStatusSeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterCourseStatusSeeder(cfg gormseeder.SeederConfiguration) *MasterCourseStatusSeeder {
	return &MasterCourseStatusSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitCourseStatusData() []entity.MasterCourseStatus {
	return []entity.MasterCourseStatus{
		{
			ID:   1,
			Name: "not_started",
		},
		{
			ID:   2,
			Name: "ongoing",
		},
		{
			ID:   3,
			Name: "completed",
		},
	}
}

func (s *MasterCourseStatusSeeder) Seed(db *gorm.DB) error {
	courseStatuses := InitCourseStatusData()

	return db.CreateInBatches(courseStatuses, len(courseStatuses)).Error
}

func (s *MasterCourseStatusSeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterCourseStatus{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
