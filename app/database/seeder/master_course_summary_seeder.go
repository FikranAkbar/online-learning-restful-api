package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/model/entity"
)

type MasterCourseSummarySeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterCourseSummarySeeder(cfg gormseeder.SeederConfiguration) *MasterCourseSummarySeeder {
	return &MasterCourseSummarySeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitCourseSummaryData() []entity.MasterCourseSummary {
	return []entity.MasterCourseSummary{
		{
			Model:    gorm.Model{ID: 1},
			CourseId: 1,
		},
		{
			Model:    gorm.Model{ID: 2},
			CourseId: 2,
		},
		{
			Model:    gorm.Model{ID: 3},
			CourseId: 3,
		},
		{
			Model:    gorm.Model{ID: 4},
			CourseId: 4,
		},
		{
			Model:    gorm.Model{ID: 5},
			CourseId: 5,
		},
		{
			Model:    gorm.Model{ID: 6},
			CourseId: 6,
		},
		{
			Model:    gorm.Model{ID: 7},
			CourseId: 7,
		},
		{
			Model:    gorm.Model{ID: 8},
			CourseId: 8,
		},
		{
			Model:    gorm.Model{ID: 9},
			CourseId: 9,
		},
		{
			Model:    gorm.Model{ID: 10},
			CourseId: 10,
		},
		{
			Model:    gorm.Model{ID: 11},
			CourseId: 11,
		},
		{
			Model:    gorm.Model{ID: 12},
			CourseId: 12,
		},
	}
}

func (s *MasterCourseSummarySeeder) Seed(db *gorm.DB) error {
	courseSummaries := InitCourseSummaryData()

	return db.CreateInBatches(courseSummaries, len(courseSummaries)).Error
}

func (s *MasterCourseSummarySeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterCourseSummary{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
