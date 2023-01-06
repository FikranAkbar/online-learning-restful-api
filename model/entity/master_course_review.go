package entity

import "online-learning-restful-api/core"

type MasterCourseReview struct {
	core.EntityModel
	CourseId   uint    `gorm:"not null"`
	UserId     uint    `gorm:"not null"`
	ReviewDesc string  `gorm:"not null"`
	ReviewRate float32 `gorm:"not null"`
}

func (MasterCourseReview) TableName() string {
	return "master_course_review"
}
