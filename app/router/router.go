package router

import (
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/controller/authentication_controller"
)

func InitRoutes(controller authentication_controller.AuthenticationController) *echo.Echo {
	e := echo.New()

	apiGroup := e.Group("/api")
	usersRouteGroup := apiGroup.Group("/users")
	usersRouteGroup.POST("/login", controller.LoginUserWithEmailPassword)

	return e
}
