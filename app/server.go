package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"online-learning-restful-api/app/router"
	"online-learning-restful-api/controller/authentication_controller"
	"online-learning-restful-api/exception"
)

func InitServerWithEcho(authenticationController authentication_controller.AuthenticationController) *echo.Echo {
	e := echo.New()

	router.InitRoutes(authenticationController, e)

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use()

	e.HTTPErrorHandler = exception.ErrorHandler

	return e
}
