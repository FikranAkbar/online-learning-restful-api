package course_service

import (
	"context"
	"online-learning-restful-api/model/web/course"
)

type CourseCategoryService interface {
	GetAllCategories(ctx context.Context) []course.CategoryCourseResponse
	GetCategory(ctx context.Context) course.CategoryCourseResponse
	GetCoursesByCategoryId(ctx context.Context, categoryId uint) []course.ShortCourseResponse
}
