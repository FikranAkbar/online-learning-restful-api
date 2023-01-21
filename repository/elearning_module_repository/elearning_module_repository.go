package elearning_module_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type ElearningModuleRepository interface {
	GetOverviewElearningModulesByCourseId(ctx context.Context, db *gorm.DB, courseId uint) ([]domain.ElearningModule, error)
}
