package authentication_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"online-learning-restful-api/model/web/authentication"
	"online-learning-restful-api/repository/account_repository"
	"online-learning-restful-api/repository/user_repository"
)

type AuthenticationServiceImpl struct {
	account_repository.AccountRepository
	user_repository.UserRepository
	*gorm.DB
	*validator.Validate
}

func NewAuthenticationServiceImpl(
	accountRepository account_repository.AccountRepository,
	userRepository user_repository.UserRepository,
	DB *gorm.DB,
	validate *validator.Validate,
) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{
		AccountRepository: accountRepository,
		UserRepository:    userRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *AuthenticationServiceImpl) LoginUserByEmailPassword(ctx context.Context, request authentication.UserLoginRequest) {
	//TODO implement me
	panic("implement me")
}