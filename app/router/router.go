package router

import (
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/controller/authentication_controller"
)

func InitRoutes(controller authentication_controller.AuthenticationController, e *echo.Echo) {
	apiGroup := e.Group("/api")
	usersRouteGroup := apiGroup.Group("/users")
	usersRouteGroup.POST("/login", controller.LoginUserWithEmailPassword).Name = "Login with email and password"
}
