package entity

import (
	"time"
)

type MasterQuiz struct {
	ID           uint `gorm:"column:id;primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ModuleId     uint   `gorm:"column:module_id;not null"`
	QuizQuestion string `gorm:"column:quiz_question;not null"`
	IsPublished  bool   `gorm:"column:is_published;default:false"`
}

func (MasterQuiz) TableName() string {
	return "master_quiz"
}
