package entity

import "online-learning-restful-api/core"

type MasterCourseSummary struct {
	core.EntityModel `gorm:"embedded"`
	CourseId         uint         `gorm:"not null"`
	MasterCourse     MasterCourse `gorm:"foreignKey:CourseId"`
	DocURL           string       `gorm:"not null"`
}

func (MasterCourseSummary) TableName() string {
	return "master_course_summary"
}
