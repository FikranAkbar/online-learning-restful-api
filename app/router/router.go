package router

import (
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/controller/authentication_controller"
	"online-learning-restful-api/controller/course_controller"
)

func InitRoutes(
	authenticationController authentication_controller.AuthenticationController,
	courseCategoryController course_controller.CourseCategoryController,
	e *echo.Echo,
) {
	apiGroup := e.Group("/api")

	// users route
	publicUserRouteGroup := apiGroup.Group("/users")
	protectedUserRouteGroup := apiGroup.Group("/users")

	// courses route
	publicCourseRouteGroup := apiGroup.Group("/courses")

	protectedRouteGroups := []*echo.Group{
		protectedUserRouteGroup,
	}

	for _, routeGroup := range protectedRouteGroups {
		routeGroup.Use(middleware.JWTAuthorization)
	}

	// authentication route
	publicUserRouteGroup.POST("/login", authenticationController.LoginUserWithEmailPassword).Name = "Login with email and password"
	publicUserRouteGroup.POST("/register", authenticationController.RegisterUserWithEmailPassword).Name = "Register user"
	protectedUserRouteGroup.POST("/logout", authenticationController.LogoutUser).Name = "Logout user"

	// course category route
	publicCourseRouteGroup.GET("/categories", courseCategoryController.GetAllCourseCategories).Name = "Get all course's categories"
	publicCourseRouteGroup.GET("/categories/:categoryId", courseCategoryController.GetCoursesByCategoryId).Name = "Get courses by category id"
}
