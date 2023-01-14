package user_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindUserById(ctx context.Context, db *gorm.DB, userId int) (domain.User, error) {
	var userEntity entity.MasterUser
	err := db.WithContext(ctx).First(&userEntity, "ID = ?", userId).Error

	if exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Sprintf("User data with id %v not found", userId)
		return domain.User{}, errors.New(logError)
	}

	return domain.User{
		Id:   int(userEntity.ID),
		Name: userEntity.Name,
	}, nil
}

func (repository *UserRepositoryImpl) CreateUserData(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error) {
	userEntity := entity.MasterUser{
		Model:  gorm.Model{ID: uint(user.Id)},
		Name:   user.Name,
		Gender: user.Gender,
		BirthDate: sql.NullTime{
			Time: user.BirthDate,
		},
	}
	err := db.WithContext(ctx).Create(&userEntity).Error
	helper.PanicIfError(err)

	return domain.User{
		Id:        int(userEntity.ID),
		Name:      userEntity.Name,
		BirthDate: userEntity.BirthDate.Time,
		Gender:    userEntity.Gender,
	}, nil
}
