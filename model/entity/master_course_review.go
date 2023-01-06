package entity

import "online-learning-restful-api/core"

type MasterCourseReview struct {
	core.EntityModel `gorm:"embedded"`
	CourseId         uint         `gorm:"not null"`
	MasterCourse     MasterCourse `gorm:"foreignKey:CourseId"`
	UserId           uint         `gorm:"not null"`
	MasterUser       MasterUser   `gorm:"foreignKey:UserId"`
	ReviewDesc       string       `gorm:"not null"`
	ReviewRate       float32      `gorm:"not null"`
}

func (MasterCourseReview) TableName() string {
	return "master_course_review"
}
