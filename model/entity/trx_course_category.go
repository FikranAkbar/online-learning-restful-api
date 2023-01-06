package entity

import "online-learning-restful-api/core"

type TrxCourseCategory struct {
	core.EntityModel `gorm:"embedded"`
	CourseId         uint           `gorm:"not null"`
	MasterCourse     MasterCourse   `gorm:"foreignKey:CourseId"`
	CategoryId       uint           `gorm:"not null"`
	MasterCategory   MasterCategory `gorm:"foreignKey:CategoryId"`
}

func (TrxCourseCategory) TableName() string {
	return "trx_course_category"
}
