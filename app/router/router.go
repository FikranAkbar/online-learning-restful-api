package router

import (
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/controller/authentication_controller"
	"online-learning-restful-api/controller/course_controller"
	"online-learning-restful-api/controller/webinar_session_controller"
)

var (
	HostURL     = "http://localhost:8000/api"
	HostURLTest = "http://localhost:8001/api"
)

// Users Routes
var (
	UsersURLPath    = "/users"
	LoginURLPath    = "/login"
	LogoutURLPath   = "/logout"
	RegisterURLPath = "/register"
)

// Course Routes
var (
	CourseURLPath             = "/courses"
	CourseIdPath              = "/:courseId"
	CourseProgressionsURLPath = "/progressions"
	CategoriesURLPath         = "/categories"
	CategoryIdPath            = "/:categoryId"
	PopularURLPath            = "/popular"
	ReviewsURLPath            = "/reviews"
	OverviewURLPath           = "/overview"
	WebinarSessionsURLPath    = "/webinar-sessions"
	WebinarSessionIdPath      = "/:webinarSessionId"
	LearnURLPath              = "/learn"
)

func InitRoutes(
	authenticationController authentication_controller.AuthenticationController,
	courseCategoryController course_controller.CourseCategoryController,
	popularCourseController course_controller.PopularCourseController,
	detailCourseController course_controller.DetailCourseController,
	courseReviewController course_controller.CourseReviewController,
	webinarSessionController webinar_session_controller.WebinarSessionController,
	e *echo.Echo,
) {
	apiGroup := e.Group("/api")

	// users route
	publicUserRouteGroup := apiGroup.Group(UsersURLPath)
	protectedUserRouteGroup := apiGroup.Group(UsersURLPath)

	// courses route
	publicCourseRouteGroup := apiGroup.Group(CourseURLPath)
	protectedCourseRouteGroup := apiGroup.Group(CourseURLPath)

	protectedRouteGroups := []*echo.Group{
		protectedUserRouteGroup,
		protectedCourseRouteGroup,
	}

	for _, routeGroup := range protectedRouteGroups {
		routeGroup.Use(middleware.JWTAuthorization)
	}

	// authentication route
	publicUserRouteGroup.POST(
		LoginURLPath,
		authenticationController.LoginUserWithEmailPassword,
	).Name = "Login with email and password"
	publicUserRouteGroup.POST(
		RegisterURLPath,
		authenticationController.RegisterUserWithEmailPassword,
	).Name = "Register user"
	protectedUserRouteGroup.POST(
		LogoutURLPath,
		authenticationController.LogoutUser,
	).Name = "Logout user"

	// course category route
	publicCourseRouteGroup.GET(
		CategoriesURLPath,
		courseCategoryController.GetAllCourseCategories,
	).Name = "Get all course's categories"
	publicCourseRouteGroup.GET(
		CategoriesURLPath+CategoryIdPath,
		courseCategoryController.GetCoursesByCategoryId,
	).Name = "Get courses by category id"

	// popular course route
	publicCourseRouteGroup.GET(
		PopularURLPath,
		popularCourseController.GetPopularCourses,
	).Name = "Get Popular Courses"

	// course route
	publicCourseRouteGroup.GET(
		"",
		detailCourseController.GetCoursesByKeyword,
	).Name = "Get Courses By Keyword"
	publicCourseRouteGroup.GET(
		CourseIdPath,
		detailCourseController.GetDetailCourseByCourseId,
	).Name = "Get Detail Course By Course Id and User Id"
	protectedCourseRouteGroup.GET(
		CourseIdPath+CourseProgressionsURLPath,
		detailCourseController.GetUserCourseProgressionByCourseId,
	).Name = "Get User Course Progression By Course Id"
	publicCourseRouteGroup.GET(
		CourseIdPath+ReviewsURLPath,
		courseReviewController.GetCourseReviewsByCourseId,
	).Name = "Get Course Review By Course Id"

	// webinar session route
	publicCourseRouteGroup.GET(
		CourseIdPath+OverviewURLPath+WebinarSessionsURLPath,
		webinarSessionController.GetOverviewWebinarSessionsByCourseId,
	).Name = "Get overview of webinar sessions by course id"
	protectedCourseRouteGroup.GET(
		CourseIdPath+LearnURLPath+WebinarSessionsURLPath+WebinarSessionIdPath,
		webinarSessionController.GetDetailWebinarSessionsByWebinarSessionId,
	).Name = "Get detail of webinar sessions by course id"
}
