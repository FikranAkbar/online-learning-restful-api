package authentication_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
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

func (service *AuthenticationServiceImpl) LoginUserByEmailPassword(ctx context.Context, request authentication.UserLoginRequest) authentication.UserLoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	account, err := service.AccountRepository.FindUserByEmail(ctx, service.DB, request.Email)
	helper.PanicIfError(err)

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password))
	helper.PanicIfError(err)

	user, err := service.UserRepository.FindUserById(ctx, service.DB, account.Id)
	helper.PanicIfError(err)

	return authentication.UserLoginResponse{
		Email: account.Email,
		Name:  user.Name,
		Token: helper.GenerateJWT(uint(account.Id)),
	}
}
