package router

import (
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/controller/authentication_controller"
	"online-learning-restful-api/controller/course_controller"
	"online-learning-restful-api/controller/elearning_module_controller"
	"online-learning-restful-api/controller/expert_controller"
	"online-learning-restful-api/controller/industry_insight_controller"
	"online-learning-restful-api/controller/quiz_controller"
	"online-learning-restful-api/controller/user_controller"
	"online-learning-restful-api/controller/webinar_session_controller"
)

var (
	HostURL     = "http://localhost:8000/api"
	HostURLTest = "http://localhost:8001/api"
)

// Users URL
var (
	UsersURLPath       = "/users"
	LoginURLPath       = "/login"
	LogoutURLPath      = "/logout"
	RegisterURLPath    = "/register"
	EditProfileURLPath = "/edit-profile"
	ProvincesURLPath   = "/provinces"
	ProvinceIdPath     = "/:provinceId"
	CitiesURLPath      = "/cities"
)

// Course URL
var (
	CourseURLPath               = "/courses"
	CourseIdPath                = "/:courseId"
	CourseProgressionsURLPath   = "/progressions"
	CategoriesURLPath           = "/categories"
	CategoryIdPath              = "/:categoryId"
	PopularURLPath              = "/popular"
	ReviewsURLPath              = "/reviews"
	OverviewURLPath             = "/overview"
	WebinarSessionsURLPath      = "/webinar-sessions"
	WebinarSessionIdPath        = "/:webinarSessionId"
	ElearningModuleURLPath      = "/modules"
	ElearningModuleIdPath       = "/:moduleId"
	LearnURLPath                = "/learn"
	SaveVideoProgressionURLPath = "/save-video-progressions"
	QuizAnswersURLPath          = "/quiz-answers"
	ComingSoonURLPath           = "/coming-soon"
	SummaryURLPath              = "/summary"
)

// Industry Insight URL
var (
	IndustryInsightsURLPath = "/industry-insights"
	IndustryInsightIdPath   = "/:industryInsightId"
)

// Expert URL
var (
	ExpertsURLPath = "/experts"
	ExpertIdPath   = "/:expertId"
)

func InitRoutes(
	authenticationController authentication_controller.AuthenticationController,
	courseCategoryController course_controller.CourseCategoryController,
	popularCourseController course_controller.PopularCourseController,
	detailCourseController course_controller.DetailCourseController,
	courseReviewController course_controller.CourseReviewController,
	webinarSessionController webinar_session_controller.WebinarSessionController,
	elearningModuleController elearning_module_controller.ElearningModuleController,
	quizController quiz_controller.QuizController,
	courseComingSoonController course_controller.ComingSoonCourseController,
	courseSummaryController course_controller.CourseSummaryController,
	industryInsightController industry_insight_controller.IndustryInsightController,
	expertController expert_controller.ExpertController,
	userController user_controller.UserController,
	e *echo.Echo,
) {
	apiGroup := e.Group("/api")

	// users group route
	publicUserRouteGroup := apiGroup.Group(UsersURLPath)
	protectedUserRouteGroup := apiGroup.Group(UsersURLPath)

	// courses group route
	publicCourseRouteGroup := apiGroup.Group(CourseURLPath)
	protectedCourseRouteGroup := apiGroup.Group(CourseURLPath)

	// industry insights group route
	publicIndustryInsightsGroup := apiGroup.Group(IndustryInsightsURLPath)

	// experts group route
	publicExpertsGroup := apiGroup.Group(ExpertsURLPath)

	// protected route group
	for _, routeGroup := range []*echo.Group{
		protectedUserRouteGroup,
		protectedCourseRouteGroup,
	} {
		routeGroup.Use(middleware.JWTAuthorization)
	}

	// authentication api route
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

	// course category api route
	publicCourseRouteGroup.GET(
		CategoriesURLPath,
		courseCategoryController.GetAllCourseCategories,
	).Name = "Get all course's categories"
	publicCourseRouteGroup.GET(
		CategoriesURLPath+CategoryIdPath,
		courseCategoryController.GetCoursesByCategoryId,
	).Name = "Get courses by category id"

	// popular course api route
	publicCourseRouteGroup.GET(
		PopularURLPath,
		popularCourseController.GetPopularCourses,
	).Name = "Get Popular Courses"

	// course api route
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

	// webinar session api route
	publicCourseRouteGroup.GET(
		CourseIdPath+OverviewURLPath+WebinarSessionsURLPath,
		webinarSessionController.GetOverviewWebinarSessionsByCourseId,
	).Name = "Get overview of webinar sessions by course id"
	protectedCourseRouteGroup.GET(
		CourseIdPath+LearnURLPath+WebinarSessionsURLPath+WebinarSessionIdPath,
		webinarSessionController.GetDetailWebinarSessionsByWebinarSessionId,
	).Name = "Get detail of webinar session by webinar session id"

	// elearning module api route
	publicCourseRouteGroup.GET(
		CourseIdPath+OverviewURLPath+ElearningModuleURLPath,
		elearningModuleController.GetOverviewElearningModulesByCourseId,
	).Name = "Get overview of elearning modules by course id"
	protectedCourseRouteGroup.GET(
		CourseIdPath+LearnURLPath+ElearningModuleURLPath+ElearningModuleIdPath,
		elearningModuleController.GetDetailElearningModuleByElearningModuleId,
	).Name = "Get detail of elearning module by elearning module id"
	protectedCourseRouteGroup.POST(
		CourseIdPath+LearnURLPath+ElearningModuleURLPath+ElearningModuleIdPath+SaveVideoProgressionURLPath,
		elearningModuleController.SaveVideoProgressionInModule,
	).Name = "Save video progression in Module"

	// quiz api route
	protectedCourseRouteGroup.GET(
		CourseIdPath+LearnURLPath+ElearningModuleURLPath+ElearningModuleIdPath+QuizAnswersURLPath,
		quizController.GetQuizAnswersByModuleId,
	).Name = "Get quiz answers by module id"
	protectedCourseRouteGroup.POST(
		CourseIdPath+LearnURLPath+ElearningModuleURLPath+ElearningModuleIdPath+QuizAnswersURLPath,
		quizController.CreateNewQuizAnswer,
	).Name = "Create new quiz answer"

	// coming soon course api route
	publicCourseRouteGroup.GET(
		ComingSoonURLPath,
		courseComingSoonController.GetComingSoonCourses,
	).Name = "Get coming soon courses"

	// course summary api route
	protectedCourseRouteGroup.GET(
		CourseIdPath+SummaryURLPath,
		courseSummaryController.GetCourseSummary,
	).Name = "Get course summary"

	// industry insight api route
	publicIndustryInsightsGroup.GET(
		"",
		industryInsightController.GetIndustryInsights,
	).Name = "Get industry insights"
	publicIndustryInsightsGroup.GET(
		IndustryInsightIdPath,
		industryInsightController.GetIndustryInsightById,
	).Name = "Get industry insights By Id"

	// expert api route
	publicExpertsGroup.GET(
		ExpertIdPath,
		expertController.GetExpertDetailById,
	).Name = "Get expert detail by expert id"
	publicExpertsGroup.GET(
		ExpertIdPath+CourseURLPath,
		expertController.GetExpertCoursesById,
	).Name = "Get expert's courses by expert id"

	// user api route
	protectedUserRouteGroup.GET(
		CourseURLPath,
		userController.GetUserCourses,
	).Name = "Get user courses"
	protectedUserRouteGroup.GET(
		EditProfileURLPath+ProvincesURLPath,
		userController.GetProvinces,
	).Name = "Get Provinces Data"
	protectedUserRouteGroup.GET(
		EditProfileURLPath+ProvincesURLPath+ProvinceIdPath+CitiesURLPath,
		userController.GetCitiesByProvinceId,
	).Name = "Get Cities By Province Id"
}
