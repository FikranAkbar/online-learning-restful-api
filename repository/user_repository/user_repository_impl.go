package user_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindUserById(ctx context.Context, tx *gorm.Tx, userId int) entity.MasterUser {
	//TODO implement me
	panic("implement me")
}
