package router

import (
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/controller/authentication_controller"
)

func InitRoutes(controller authentication_controller.AuthenticationController, e *echo.Echo) {
	apiGroup := e.Group("/api")

	publicUserRouteGroup := apiGroup.Group("/users")
	protectedUserRouteGroup := apiGroup.Group("/users")

	protectedRouteGroups := []*echo.Group{
		protectedUserRouteGroup,
	}

	for _, routeGroup := range protectedRouteGroups {
		//routeGroup.Use(echojwt.WithConfig(echojwt.Config{SigningKey: helper.JwtSignatureKey}))
		routeGroup.Use(middleware.JWTAuthorization)
	}

	publicUserRouteGroup.POST("/login", controller.LoginUserWithEmailPassword).Name = "Login with email and password"
	publicUserRouteGroup.POST("/register", controller.RegisterUserWithEmailPassword).Name = "Register user"
	protectedUserRouteGroup.POST("/logout", controller.LogoutUser).Name = "Logout user"
}
