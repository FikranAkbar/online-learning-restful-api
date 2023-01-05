package domain

import "online-learning-restful-api/core"

type MasterCourseSummary struct {
	core.DomainModel
	CourseId uint   `gorm:"not null"`
	DocURL   string `gorm:"not null"`
}
