package course_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type CourseRepository interface {
	GetPopularCourses(ctx context.Context, db *gorm.DB) ([]domain.Course, error)
	GetCoursesByKeyword(ctx context.Context, db *gorm.DB, keyword string) ([]domain.Course, error)
	GetCourseDetailByCourseId(ctx context.Context, db *gorm.DB, courseId uint, userId *uint) (domain.Course, domain.Expert, error)
	GetUserCourseProgressionByCourseId(ctx context.Context, db *gorm.DB, courseId uint) (domain.UserCourse, error)
	GetCourseReviewsByCourseId(ctx context.Context, db *gorm.DB, courseId uint) ([]domain.CourseReview, error)
	GetComingSoonCourses(ctx context.Context, db *gorm.DB) ([]domain.Course, error)
}
