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

func (service *CourseCategoryServiceImpl) GetAllCategories(ctx context.Context) []course.CategoryCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	categories, err := service.CategoryRepository.GetAllCategories(ctx, tx)
	helper.PanicIfError(err)

	var categoriesResponse []course.CategoryCourseResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, course.CategoryCourseResponse{
			Id:   category.Id,
			Name: category.Name,
		})
	}

	return categoriesResponse
}

func (service *CourseCategoryServiceImpl) GetCategory(ctx context.Context) course.CategoryCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.GetCategory(ctx, tx)
	helper.PanicIfError(err)

	return course.CategoryCourseResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func (service *CourseCategoryServiceImpl) GetCoursesByCategoryId(ctx context.Context, categoryId uint) []course.ShortCourseResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	courses, err := service.CategoryRepository.GetCoursesByCategoryId(ctx, tx, categoryId)
	helper.PanicIfError(err)

	var coursesResponse []course.ShortCourseResponse
	for _, courseData := range courses {
		coursesResponse = append(coursesResponse, course.ShortCourseResponse{
			Id:          courseData.Id,
			Name:        courseData.Name,
			PhotoUrl:    courseData.PhotoUrl,
			AverageRate: courseData.AverageRate,
			ExpertName:  courseData.ExpertName,
		})
	}

	return coursesResponse
}
