package domain

import "online-learning-restful-api/core"

type MasterQuiz struct {
	core.DomainModel
	ModuleId     uint `gorm:"not null"`
	QuizQuestion string
	IsPublished  bool `gorm:"default:false"`
}
