package course_service

import (
	"context"
	"online-learning-restful-api/model/web/course"
)

type CourseCategoryService interface {
	GetAllCategories(ctx context.Context) []course.CategoryResponse
	GetCategory(ctx context.Context) course.CategoryResponse
}
