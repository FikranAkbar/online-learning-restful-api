package router

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/controller/authentication_controller"
	"online-learning-restful-api/helper"
)

func InitRoutes(controller authentication_controller.AuthenticationController, e *echo.Echo) {
	apiGroup := e.Group("/api")

	publicRoute := apiGroup
	protectedRoute := apiGroup
	protectedRoute.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: helper.JwtSecret,
	}))

	userPublicRoutes := publicRoute.Group("/users")

	userPublicRoutes.POST("/login", controller.LoginUserWithEmailPassword).Name = "Login with email and password"
	userPublicRoutes.POST("/register", controller.RegisterUserWithEmailPassword).Name = "Register User"
}
