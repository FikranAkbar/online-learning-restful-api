package category_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type CategoryRepository interface {
	GetAllCategories(ctx context.Context, db *gorm.DB) ([]domain.Category, error)
}
