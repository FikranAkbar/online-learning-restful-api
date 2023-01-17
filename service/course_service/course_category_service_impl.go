package course_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/repository/category_repository"
)

type CourseCategoryServiceImpl struct {
	category_repository.CategoryRepository
	*gorm.DB
}

func NewCourseCategoryServiceImpl(repository category_repository.CategoryRepository, db *gorm.DB) *CourseCategoryServiceImpl {
	return &CourseCategoryServiceImpl{
		CategoryRepository: repository,
		DB:                 db,
	}
}

func (service CourseCategoryServiceImpl) GetAllCategories(ctx context.Context) []course.CategoryResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	categories, err := service.CategoryRepository.GetAllCategories(ctx, tx)
	helper.PanicIfError(err)

	var categoriesResponse []course.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, course.CategoryResponse{
			Id:   category.Id,
			Name: category.Name,
		})
	}

	return categoriesResponse
}

func (service CourseCategoryServiceImpl) GetCategory(ctx context.Context) course.CategoryResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.GetCategory(ctx, tx)
	helper.PanicIfError(err)

	return course.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
