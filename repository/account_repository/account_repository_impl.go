package account_repository

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
)

type AccountRepositoryImpl struct {
}

func NewAccountRepositoryImpl() *AccountRepositoryImpl {
	return &AccountRepositoryImpl{}
}

func (repository *AccountRepositoryImpl) FindUserByEmail(ctx context.Context, db *gorm.DB, email string) (domain.Account, error) {
	var accountEntity entity.MasterAccount
	err := db.WithContext(ctx).Preload("MasterUserType").First(&accountEntity, "email = ?", email).Error

	if exception.CheckErrorContains(err, exception.NotFound) {
		errorLog := fmt.Sprintf("Wrong email or password")
		return domain.Account{}, exception.GenerateHTTPError(exception.Unauthorized, errorLog)
	}

	return domain.Account{
		Id:       int(accountEntity.ID),
		Email:    accountEntity.Email,
		Password: accountEntity.Password,
		Role:     accountEntity.MasterUserType.Name,
	}, nil
}

func (repository *AccountRepositoryImpl) CreateAccountData(ctx context.Context, db *gorm.DB, account domain.Account) (domain.Account, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	roleString := account.Role
	var userType entity.MasterUserType
	err = db.WithContext(ctx).First(&userType).Where("Name = ?", roleString).Error
	helper.PanicIfError(err)

	accountEntity := entity.MasterAccount{
		Email:    account.Email,
		Password: string(hashedPassword),
		Role:     userType.ID,
	}
	err = db.WithContext(ctx).Create(&accountEntity).Error
	if exception.CheckErrorContains(err, exception.Duplicate) {
		logError := fmt.Sprintf("Account with email %v already exist", accountEntity.Email)
		return domain.Account{}, exception.GenerateHTTPError(exception.BadRequest, logError)
	}

	return domain.Account{
		Id:       int(accountEntity.ID),
		Email:    accountEntity.Email,
		Password: accountEntity.Password,
		Role:     roleString,
	}, nil
}
