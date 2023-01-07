package entity

import "online-learning-restful-api/core"

type TrxCourseCategory struct {
	core.EntityModel `gorm:"embedded"`
	CourseId         uint           `gorm:"column:course_id;not null"`
	MasterCourse     MasterCourse   `gorm:"foreignKey:CourseId"`
	CategoryId       uint           `gorm:"column:category_id;not null"`
	MasterCategory   MasterCategory `gorm:"foreignKey:CategoryId"`
}

func (TrxCourseCategory) TableName() string {
	return "trx_course_category"
}
