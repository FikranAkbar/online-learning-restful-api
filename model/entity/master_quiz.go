package entity

import "online-learning-restful-api/core"

type MasterQuiz struct {
	core.EntityModel      `gorm:"embedded"`
	ModuleId              uint                  `gorm:"not null"`
	MasterElearningModule MasterElearningModule `gorm:"foreignKey:ModuleId"`
	QuizQuestion          string                `gorm:"not null"`
	IsPublished           bool                  `gorm:"default:false"`
}

func (MasterQuiz) TableName() string {
	return "master_quiz"
}
