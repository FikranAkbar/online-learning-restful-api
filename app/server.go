package app

import (
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/app/router"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/controller/authentication_controller"
	"online-learning-restful-api/controller/course_controller"
	"online-learning-restful-api/exception"
)

func InitServerWithEcho(
	authenticationController authentication_controller.AuthenticationController,
	courseCategoryController course_controller.CourseCategoryController,
	coursePopularController course_controller.CoursePopularController,
	courseController course_controller.CourseController,
) *echo.Echo {
	e := echo.New()

	router.InitRoutes(
		authenticationController,
		courseCategoryController,
		coursePopularController,
		courseController,
		e)
	middleware.InitMiddleware(e)

	e.HTTPErrorHandler = exception.ErrorHandler

	return e
}
