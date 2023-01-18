package category_repository

import (
	"context"
	"fmt"
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

func (repository *CategoryRepositoryImpl) GetCategory(ctx context.Context, db *gorm.DB) (domain.Category, error) {
	var categoryEntity entity.MasterCategory
	err := db.WithContext(ctx).First(&categoryEntity).Error
	if err != nil {
		return domain.Category{}, err
	}

	return domain.Category{
		Id:   categoryEntity.ID,
		Name: categoryEntity.CategoryName,
	}, nil
}

func (repository *CategoryRepositoryImpl) GetCoursesByCategoryId(ctx context.Context, db *gorm.DB, categoryId uint) ([]domain.Course, error) {
	var categoryEntity entity.MasterCategory
	err := db.WithContext(ctx).Where("ID = ?", categoryId).Preload("Courses").First(&categoryEntity).Error
	if err != nil {
		return []domain.Course{}, err
	}

	courseEntities := categoryEntity.Courses
	err = db.WithContext(ctx).Preload("Expert").Error
	if err != nil {
		return []domain.Course{}, err
	}

	fmt.Println(courseEntities)

	var courses []domain.Course
	for _, courseEntity := range courseEntities {
		courses = append(courses, domain.Course{
			Id:          courseEntity.ID,
			Name:        courseEntity.Name,
			PhotoUrl:    courseEntity.PhotoURL.String,
			AverageRate: courseEntity.AverageRate,
			ExpertName:  courseEntity.Expert.Name,
		})
	}

	return courses, nil
}
