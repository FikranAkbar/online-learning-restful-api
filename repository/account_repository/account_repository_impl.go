package account_repository

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type AccountRepositoryImpl struct {
}

func NewAccountRepositoryImpl() *AccountRepositoryImpl {
	return &AccountRepositoryImpl{}
}

func (repository AccountRepositoryImpl) FindUserByEmail(ctx context.Context, tx *gorm.Tx, email string) entity.MasterAccount {
	//TODO implement me
	panic("implement me")
}
