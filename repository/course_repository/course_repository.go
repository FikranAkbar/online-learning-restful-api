package course_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type CourseRepository interface {
	GetPopularCourses(ctx context.Context, db *gorm.DB) ([]domain.Course, error)
}
