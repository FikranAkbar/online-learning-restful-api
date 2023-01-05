package domain

import "online-learning-restful-api/core"

type MasterCourseReview struct {
	core.DomainModel
	CourseId   uint    `gorm:"not null"`
	UserId     uint    `gorm:"not null"`
	ReviewDesc string  `gorm:"not null"`
	ReviewRate float32 `gorm:"not null"`
}
