package user_repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/model/domain"
	"strings"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindUserById(ctx context.Context, db *gorm.DB, userId int) (*domain.User, error) {
	var userEntity entity.MasterUser
	err := db.WithContext(ctx).First(&userEntity, "ID = ?", userId).Error

	if strings.Contains(err.Error(), exception.NotFoundError) {
		logError := fmt.Sprintf("User data with id %v not found", userId)
		return nil, errors.New(logError)
	}

	return &domain.User{
		Id:   int(userEntity.ID),
		Name: userEntity.Name,
	}, nil
}
