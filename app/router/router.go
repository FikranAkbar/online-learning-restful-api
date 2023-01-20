package router

import (
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/controller/authentication_controller"
	"online-learning-restful-api/controller/course_controller"
)

var (
	HostURL     = "http://localhost:8000/api"
	HostURLTest = "http://localhost:8001/api"
)

// Users Routes
var (
	UsersAPIRoute    = "/users"
	LoginAPIRoute    = "/login"
	LogoutAPIRoute   = "/logout"
	RegisterAPIRoute = "/register"
)

// Course Categories Routes
var (
	CoursesAPIRoute            = "/courses"
	CourseIdAPIRoute           = "/:courseId"
	CourseProgressionsAPIRoute = "/progressions"
	CategoriesAPIRoute         = "/categories"
	CategoryIdAPIRoute         = "/:categoryId"
	PopularAPIRoute            = "/popular"
	ReviewAPIRoute             = "/reviews"
)

func InitRoutes(
	authenticationController authentication_controller.AuthenticationController,
	courseCategoryController course_controller.CourseCategoryController,
	popularCourseController course_controller.PopularCourseController,
	detailCourseController course_controller.DetailCourseController,
	courseReviewController course_controller.CourseReviewController,
	e *echo.Echo,
) {
	apiGroup := e.Group("/api")

	// users route
	publicUserRouteGroup := apiGroup.Group(UsersAPIRoute)
	protectedUserRouteGroup := apiGroup.Group(UsersAPIRoute)

	// courses route
	publicCourseRouteGroup := apiGroup.Group(CoursesAPIRoute)
	protectedCourseRouteGroup := apiGroup.Group(CoursesAPIRoute)

	protectedRouteGroups := []*echo.Group{
		protectedUserRouteGroup,
		protectedCourseRouteGroup,
	}

	for _, routeGroup := range protectedRouteGroups {
		routeGroup.Use(middleware.JWTAuthorization)
	}

	// authentication route
	publicUserRouteGroup.POST(LoginAPIRoute, authenticationController.LoginUserWithEmailPassword).Name = "Login with email and password"
	publicUserRouteGroup.POST(RegisterAPIRoute, authenticationController.RegisterUserWithEmailPassword).Name = "Register user"
	protectedUserRouteGroup.POST(LogoutAPIRoute, authenticationController.LogoutUser).Name = "Logout user"

	// course category route
	publicCourseRouteGroup.GET(CategoriesAPIRoute, courseCategoryController.GetAllCourseCategories).Name = "Get all course's categories"
	publicCourseRouteGroup.GET(CategoriesAPIRoute+CategoryIdAPIRoute, courseCategoryController.GetCoursesByCategoryId).Name = "Get courses by category id"

	// popular course route
	publicCourseRouteGroup.GET(PopularAPIRoute, popularCourseController.GetPopularCourses).Name = "Get Popular Courses"

	// course route
	publicCourseRouteGroup.GET("", detailCourseController.GetCoursesByKeyword).Name = "Get Courses By Keyword"
	publicCourseRouteGroup.GET(CourseIdAPIRoute, detailCourseController.GetDetailCourseByCourseId).Name = "Get Detail Course By Course Id and User Id"
	protectedCourseRouteGroup.GET(CourseIdAPIRoute+CourseProgressionsAPIRoute, detailCourseController.GetUserCourseProgressionByCourseId).Name = "Get User Course Progression By Course Id"
	publicCourseRouteGroup.GET(CourseIdAPIRoute+ReviewAPIRoute, courseReviewController.GetCourseReviewsByCourseId).Name = "Get Course Review By Course Id"
}
