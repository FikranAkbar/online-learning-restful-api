package entity

import "online-learning-restful-api/core"

type TrxCourseCategory struct {
	core.EntityModel
	CourseId   uint `gorm:"not null"`
	CategoryId uint `gorm:"not null"`
}

func (TrxCourseCategory) TableName() string {
	return "trx_course_category"
}
