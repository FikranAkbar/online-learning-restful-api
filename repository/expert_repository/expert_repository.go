package expert_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type ExpertRepository interface {
	GetExpertDetailById(ctx context.Context, db *gorm.DB, expertId uint) (domain.Expert, error)
	GetExpertCoursesById(ctx context.Context, db *gorm.DB, expertId uint) ([]domain.Course, error)
}
