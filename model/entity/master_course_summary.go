package entity

import "online-learning-restful-api/core"

type MasterCourseSummary struct {
	core.EntityModel
	CourseId uint   `gorm:"not null"`
	DocURL   string `gorm:"not null"`
}

func (MasterCourseSummary) TableName() string {
	return "master_course_summary"
}
