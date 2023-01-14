package user_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/model/domain"
)

type UserRepository interface {
	FindUserById(ctx context.Context, db *gorm.DB, userId int) (domain.User, error)
	CreateUserData(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error)
}
