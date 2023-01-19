package seeder

import (
	"database/sql"
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type TrxUserCourseSeeder struct {
	gormseeder.SeederAbstract
}

func NewTrxUserCourseSeeder(cfg gormseeder.SeederConfiguration) *TrxUserCourseSeeder {
	return &TrxUserCourseSeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitUserCourseData() []entity.TrxUserCourse {
	return []entity.TrxUserCourse{
		{
			UserId:               1,
			CourseId:             1,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
		{
			UserId:               1,
			CourseId:             2,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
		{
			UserId:               1,
			CourseId:             4,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
		{
			UserId:               2,
			CourseId:             6,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
		{
			UserId:               2,
			CourseId:             2,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
		{
			UserId:               2,
			CourseId:             7,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
		{
			UserId:               3,
			CourseId:             8,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
		{
			UserId:               3,
			CourseId:             3,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
		{
			UserId:               3,
			CourseId:             9,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
		{
			UserId:               4,
			CourseId:             11,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
		{
			UserId:               4,
			CourseId:             12,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
		{
			UserId:               4,
			CourseId:             7,
			LastUnlockedModule:   1,
			TotalWebinarAttended: 0,
			GraduatedAt:          sql.NullTime{},
		},
	}
}

func (s *TrxUserCourseSeeder) Seed(db *gorm.DB) error {
	userCourses := InitUserCourseData()

	return db.CreateInBatches(userCourses, len(userCourses)).Error
}

func (s *TrxUserCourseSeeder) Clear(db *gorm.DB) error {
	ent := &entity.TrxUserCourse{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}
