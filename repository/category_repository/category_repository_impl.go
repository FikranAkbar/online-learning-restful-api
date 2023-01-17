package category_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepositoryImpl() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) GetAllCategories(ctx context.Context, db *gorm.DB) ([]domain.Category, error) {
	var categoryEntities []entity.MasterCategory
	err := db.WithContext(ctx).Find(&categoryEntities).Error
	if err != nil {
		return []domain.Category{}, err
	}

	var categories []domain.Category
	for _, category := range categoryEntities {
		categories = append(categories, domain.Category{
			Id:   category.ID,
			Name: category.CategoryName,
		})
	}

	return categories, nil
}
