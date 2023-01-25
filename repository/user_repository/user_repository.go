package user_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type UserRepository interface {
	FindUserById(ctx context.Context, db *gorm.DB, userId uint) (domain.User, error)
	CreateUserData(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error)
	GetUserCourses(ctx context.Context, db *gorm.DB, courseStatus string) ([]domain.Course, error)
	GetProvinces(ctx context.Context, db *gorm.DB) ([]domain.Province, error)
	GetCitiesByProvinceId(ctx context.Context, db *gorm.DB, provinceId uint) ([]domain.City, error)
	EditUserProfile(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error)
}
