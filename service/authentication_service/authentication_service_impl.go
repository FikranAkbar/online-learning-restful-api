package authentication_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"online-learning-restful-api/model/web/authentication"
	"online-learning-restful-api/repository/account_repository"
	"online-learning-restful-api/repository/user_repository"
	"time"
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

func (service *AuthenticationServiceImpl) RegisterUserByEmailPassword(ctx context.Context, request authentication.UserRegisterRequest) authentication.UserRegisterResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	account := domain.Account{
		Email:    request.Email,
		Password: request.Password,
		Role:     "User",
	}
	account, err = service.AccountRepository.CreateAccountData(ctx, service.DB, account)
	helper.PanicIfError(err)

	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	helper.PanicIfError(err)
	user := domain.User{
		Id:        account.Id,
		Name:      request.Name,
		BirthDate: birthDate,
		Gender:    request.Gender,
	}
	user, err = service.UserRepository.CreateUserData(ctx, service.DB, user)
	helper.PanicIfError(err)

	userRegisterResponse := authentication.UserRegisterResponse{
		Id:    account.Id,
		Name:  user.Name,
		Email: account.Email,
	}

	return userRegisterResponse
}
