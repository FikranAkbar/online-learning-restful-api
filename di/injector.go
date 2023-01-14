//go:build wireinject
// +build wireinject

package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/app"
	"online-learning-restful-api/app/database"
	"online-learning-restful-api/controller/authentication_controller"
	"online-learning-restful-api/repository/account_repository"
	"online-learning-restful-api/repository/user_repository"
	"online-learning-restful-api/service/authentication_service"
)

var authenticationSet = wire.NewSet(
	account_repository.NewAccountRepositoryImpl,
	wire.Bind(new(account_repository.AccountRepository), new(*account_repository.AccountRepositoryImpl)),
	user_repository.NewUserRepositoryImpl,
	wire.Bind(new(user_repository.UserRepository), new(*user_repository.UserRepositoryImpl)),
	authentication_service.NewAuthenticationServiceImpl,
	wire.Bind(new(authentication_service.AuthenticationService), new(*authentication_service.AuthenticationServiceImpl)),
	authentication_controller.NewAuthenticationControllerImpl,
	wire.Bind(new(authentication_controller.AuthenticationController), new(*authentication_controller.AuthenticationControllerImpl)),
)

func InitializedEchoServer() *echo.Echo {
	wire.Build(
		app.InitServerWithEcho,
		database.NewDB,
		validator.New,
		authenticationSet,
	)

	return nil
}
