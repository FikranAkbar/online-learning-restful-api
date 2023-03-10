package authentication_service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/exception"
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
	if err != nil {
		panic(exception.GenerateHTTPError(exception.BadRequest, err.Error()))
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	account, err := service.AccountRepository.FindUserByEmail(ctx, tx, request.Email)
	helper.PanicIfError(err)

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password))
	if err != nil {
		errorLog := fmt.Sprintf("Wrong email or password")
		panic(exception.GenerateHTTPError(exception.Unauthorized, errorLog))
	}

	user, err := service.UserRepository.FindUserById(ctx, tx, account.Id)
	helper.PanicIfError(err)

	return authentication.UserLoginResponse{
		Email: account.Email,
		Name:  user.Name,
		Token: helper.GenerateJWT(account.Id, account.Email, user.Name),
	}
}

func (service *AuthenticationServiceImpl) RegisterUserByEmailPassword(ctx context.Context, request authentication.UserRegisterRequest) authentication.UserRegisterResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(exception.GenerateHTTPError(exception.BadRequest, err.Error()))
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	account := domain.Account{
		Email:    request.Email,
		Password: request.Password,
		Role:     "User",
	}

	account, err = service.AccountRepository.CreateAccountData(ctx, tx, account)
	helper.PanicIfError(err)

	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	helper.PanicIfError(err)
	user := domain.User{
		Id:        account.Id,
		Name:      request.Name,
		BirthDate: birthDate,
		Gender:    request.Gender,
	}
	user, err = service.UserRepository.CreateUserData(ctx, tx, user)
	helper.PanicIfError(err)

	userRegisterResponse := authentication.UserRegisterResponse{
		Id:    account.Id,
		Name:  user.Name,
		Email: account.Email,
	}

	return userRegisterResponse
}

func (service *AuthenticationServiceImpl) LogoutUser(ctx context.Context) string {
	userTokenInfo, ok := ctx.Value(middleware.ContextUserInfoKey).(middleware.UserTokenInfo)
	if !ok {
		panic(middleware.UnauthorizedErrorInfo)
	}

	return fmt.Sprintf("User %v logout success. Token valid", userTokenInfo.UserName)
}
