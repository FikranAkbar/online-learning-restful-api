//go:build wireinject
// +build wireinject

package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/app"
	"online-learning-restful-api/app/database"
	"online-learning-restful-api/controller/authentication_controller"
	"online-learning-restful-api/controller/course_controller"
	"online-learning-restful-api/controller/elearning_module_controller"
	"online-learning-restful-api/controller/webinar_session_controller"
	"online-learning-restful-api/repository/account_repository"
	"online-learning-restful-api/repository/category_repository"
	"online-learning-restful-api/repository/course_repository"
	"online-learning-restful-api/repository/elearning_module_repository"
	"online-learning-restful-api/repository/user_repository"
	"online-learning-restful-api/repository/webinar_session_repository"
	"online-learning-restful-api/service/authentication_service"
	"online-learning-restful-api/service/course_service"
	"online-learning-restful-api/service/elearning_module_service"
	"online-learning-restful-api/service/webinar_session_service"
)

var authenticationSet = wire.NewSet(
	account_repository.NewAccountRepositoryImpl,
	wire.Bind(new(account_repository.AccountRepository), new(*account_repository.AccountRepositoryImpl)),
	user_repository.NewUserRepositoryImpl,
	wire.Bind(new(user_repository.UserRepository), new(*user_repository.UserRepositoryImpl)),
	authentication_service.NewAuthenticationServiceImpl,
	wire.Bind(new(authentication_service.AuthenticationService), new(*authentication_service.AuthenticationServiceImpl)),
	authentication_controller.NewAuthenticationControllerImpl,
	wire.Bind(new(authentication_controller.AuthenticationController), new(*authentication_controller.AuthenticationControllerImpl)),
)

var courseCategorySet = wire.NewSet(
	category_repository.NewCategoryRepositoryImpl,
	wire.Bind(new(category_repository.CategoryRepository), new(*category_repository.CategoryRepositoryImpl)),
	course_service.NewCourseCategoryServiceImpl,
	wire.Bind(new(course_service.CourseCategoryService), new(*course_service.CourseCategoryServiceImpl)),
	course_controller.NewCourseCategoryControllerImpl,
	wire.Bind(new(course_controller.CourseCategoryController), new(*course_controller.CourseCategoryControllerImpl)),
)

var popularCourseSet = wire.NewSet(
	course_service.NewCoursePopularServiceImpl,
	wire.Bind(new(course_service.CoursePopularService), new(*course_service.CoursePopularServiceImpl)),
	course_controller.NewPopularCourseControllerImpl,
	wire.Bind(new(course_controller.PopularCourseController), new(*course_controller.PopularCourseControllerImpl)),
)

var courseRepositorySet = wire.NewSet(
	course_repository.NewCourseRepositoryImpl,
	wire.Bind(new(course_repository.CourseRepository), new(*course_repository.CourseRepositoryImpl)),
)

var detailCourseSet = wire.NewSet(
	course_controller.NewDetailCourseControllerImpl,
	wire.Bind(new(course_controller.DetailCourseController), new(*course_controller.DetailCourseControllerImpl)),
	course_service.NewCourseDetailServiceImpl,
	wire.Bind(new(course_service.CourseDetailService), new(*course_service.CourseDetailServiceImpl)),
)

var reviewCourseSet = wire.NewSet(
	course_controller.NewCourseReviewControllerImpl,
	wire.Bind(new(course_controller.CourseReviewController), new(*course_controller.CourseReviewControllerImpl)),
	course_service.NewCourseReviewServiceImpl,
	wire.Bind(new(course_service.CourseReviewService), new(*course_service.CourseReviewServiceImpl)),
)

var webinarSessionSet = wire.NewSet(
	webinar_session_repository.NewWebinarSessionRepositoryImpl,
	wire.Bind(new(webinar_session_repository.WebinarSessionRepository), new(*webinar_session_repository.WebinarSessionRepositoryImpl)),
	webinar_session_service.NewWebinarSessionServiceImpl,
	wire.Bind(new(webinar_session_service.WebinarSessionService), new(*webinar_session_service.WebinarSessionServiceImpl)),
	webinar_session_controller.NewWebinarSessionControllerImpl,
	wire.Bind(new(webinar_session_controller.WebinarSessionController), new(*webinar_session_controller.WebinarSessionControllerImpl)),
)

var elearningModuleSet = wire.NewSet(
	elearning_module_repository.NewElearningModuleRepositoryImpl,
	wire.Bind(new(elearning_module_repository.ElearningModuleRepository), new(*elearning_module_repository.ElearningModuleRepositoryImpl)),
	elearning_module_service.NewElearningModuleServiceImpl,
	wire.Bind(new(elearning_module_service.ElearningModuleService), new(*elearning_module_service.ElearningModuleServiceImpl)),
	elearning_module_controller.NewElearningModuleControllerImpl,
	wire.Bind(new(elearning_module_controller.ElearningModuleController), new(*elearning_module_controller.ElearningModuleControllerImpl)),
)

var completeSet = wire.NewSet(
	app.InitServerWithEcho,
	validator.New,
	authenticationSet,
	courseRepositorySet,
	courseCategorySet,
	popularCourseSet,
	detailCourseSet,
	reviewCourseSet,
	webinarSessionSet,
	elearningModuleSet,
)

func InitializedEchoServer() *echo.Echo {
	wire.Build(completeSet, database.NewDB)
	return nil
}

func InitializedEchoServerForTest() *echo.Echo {
	wire.Build(completeSet, database.NewTestDB)
	return nil
}
