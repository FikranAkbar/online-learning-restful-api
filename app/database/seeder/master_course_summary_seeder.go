package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
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
			ID:       1,
			CourseId: 1,
		},
		{
			ID:       2,
			CourseId: 2,
		},
		{
			ID:       3,
			CourseId: 3,
		},
		{
			ID:       4,
			CourseId: 4,
		},
		{
			ID:       5,
			CourseId: 5,
		},
		{
			ID:       6,
			CourseId: 6,
		},
		{
			ID:       7,
			CourseId: 7,
		},
		{
			ID:       8,
			CourseId: 8,
		},
		{
			ID:       9,
			CourseId: 9,
		},
		{
			ID:       10,
			CourseId: 10,
		},
		{
			ID:       11,
			CourseId: 11,
		},
		{
			ID:       12,
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
