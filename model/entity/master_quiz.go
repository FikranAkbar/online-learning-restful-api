package entity

import (
	"gorm.io/gorm"
)

type MasterQuiz struct {
	gorm.Model            `gorm:"embedded"`
	ModuleId              uint                  `gorm:"column:module_id;not null"`
	MasterElearningModule MasterElearningModule `gorm:"foreignKey:ModuleId"`
	QuizQuestion          string                `gorm:"column:quiz_question;not null"`
	IsPublished           bool                  `gorm:"column:is_published;default:false"`
}

func (MasterQuiz) TableName() string {
	return "master_quiz"
}
