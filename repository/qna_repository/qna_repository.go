package qna_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type QnaRepository interface {
	GetQnaQuestionsByCourseId(ctx context.Context, db *gorm.DB, courseId uint) ([]domain.QnaQuestion, []domain.User, error)
	CreateNewQnaQuestion(ctx context.Context, db *gorm.DB, courseId uint, qnaQuestion domain.QnaQuestion) (domain.QnaQuestion, error)
}
