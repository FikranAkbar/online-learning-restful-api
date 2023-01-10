package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type TrxCourseCategorySeeder struct {
	gormseeder.SeederAbstract
}

func NewTrxCourseCategorySeeder(cfg gormseeder.SeederConfiguration) *TrxCourseCategorySeeder {
	return &TrxCourseCategorySeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitCourseCategoryData() []entity.TrxCourseCategory {
	return []entity.TrxCourseCategory{
		{
			CourseId:   1,
			CategoryId: 1,
		},
		{
			CourseId:   2,
			CategoryId: 2,
		},
		{
			CourseId:   3,
			CategoryId: 3,
		},
		{
			CourseId:   4,
			CategoryId: 4,
		},
		{
			CourseId:   6,
			CategoryId: 2,
		},
		{
			CourseId:   7,
			CategoryId: 3,
		},
		{
			CourseId:   8,
			CategoryId: 4,
		},
		{
			CourseId:   9,
			CategoryId: 1,
		},
		{
			CourseId:   10,
			CategoryId: 2,
		},
		{
			CourseId:   11,
			CategoryId: 3,
		},
		{
			CourseId:   12,
			CategoryId: 4,
		},
	}
}

func (s *TrxCourseCategorySeeder) Seed(db *gorm.DB) error {
	data := InitCourseCategoryData()

	return db.CreateInBatches(data, len(data)).Error
}

func (s *TrxCourseCategorySeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxCourseCategory{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
