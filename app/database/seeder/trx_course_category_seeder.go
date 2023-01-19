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
	var entities []entity.TrxCourseCategory
	for i := 0; i < 12; i++ {
		entities = append(entities, entity.TrxCourseCategory{
			CourseId:   uint(i + 1),
			CategoryId: uint((i % 4) + 1),
		})
	}
	return entities
}

func (s *TrxCourseCategorySeeder) Seed(db *gorm.DB) error {
	data := InitCourseCategoryData()

	return db.CreateInBatches(data, len(data)).Error
}

func (s *TrxCourseCategorySeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxCourseCategory{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
