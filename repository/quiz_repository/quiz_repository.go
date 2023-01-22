package quiz_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type QuizRepository interface {
	GetQuizAnswersByModuleId(ctx context.Context, db *gorm.DB, courseId uint, elearningModuleId uint) ([]domain.QuizAnswer, error)
	CreateNewQuizAnswer(ctx context.Context, db *gorm.DB, courseId uint, elearningModuleId uint, quizAnswer domain.QuizAnswer) (domain.QuizAnswer, error)
}
