package account_repository

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

type AccountRepositoryImpl struct {
}

func NewAccountRepositoryImpl() *AccountRepositoryImpl {
	return &AccountRepositoryImpl{}
}

func (repository AccountRepositoryImpl) FindUserByEmail(ctx context.Context, db *gorm.DB, email string) (*domain.Account, error) {
	var accountEntity entity.MasterAccount
	err := db.WithContext(ctx).Preload("MasterUserType").First(&accountEntity, "email = ?", email).Error

	if strings.Contains(err.Error(), exception.NotFoundError) {
		logError := fmt.Sprint("Unauthorized. Wrong email or password")
		return nil, errors.New(logError)
	}

	return &domain.Account{
		Id:       int(accountEntity.ID),
		Email:    accountEntity.Email,
		Password: accountEntity.Password,
		Role:     accountEntity.MasterUserType.Name,
	}, nil
}
