// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/app"
	"online-learning-restful-api/app/database"
	"online-learning-restful-api/config"
	"online-learning-restful-api/controller/authentication_controller"
	"online-learning-restful-api/controller/cart_controller"
	"online-learning-restful-api/controller/course_controller"
	"online-learning-restful-api/controller/elearning_module_controller"
	"online-learning-restful-api/controller/expert_controller"
	"online-learning-restful-api/controller/industry_insight_controller"
	"online-learning-restful-api/controller/qna_controller"
	"online-learning-restful-api/controller/quiz_controller"
	"online-learning-restful-api/controller/user_controller"
	"online-learning-restful-api/controller/webinar_session_controller"
	"online-learning-restful-api/repository/account_repository"
	"online-learning-restful-api/repository/cart_repository"
	"online-learning-restful-api/repository/category_repository"
	"online-learning-restful-api/repository/course_repository"
	"online-learning-restful-api/repository/elearning_module_repository"
	"online-learning-restful-api/repository/expert_repository"
	"online-learning-restful-api/repository/industry_insight_repository"
	"online-learning-restful-api/repository/qna_repository"
	"online-learning-restful-api/repository/quiz_repository"
	"online-learning-restful-api/repository/user_repository"
	"online-learning-restful-api/repository/webinar_session_repository"
	"online-learning-restful-api/service/authentication_service"
	"online-learning-restful-api/service/cart_service"
	"online-learning-restful-api/service/course_service"
	"online-learning-restful-api/service/elearning_module_service"
	"online-learning-restful-api/service/expert_service"
	"online-learning-restful-api/service/industry_insight_service"
	"online-learning-restful-api/service/qna_service"
	"online-learning-restful-api/service/quiz_service"
	"online-learning-restful-api/service/user_service"
	"online-learning-restful-api/service/webinar_session_service"
)

// Injectors from injector.go:

func InitializedEchoServer() *echo.Echo {
	accountRepositoryImpl := account_repository.NewAccountRepositoryImpl()
	userRepositoryImpl := user_repository.NewUserRepositoryImpl()
	configConfig := config.LoadConfigFromEnv()
	db := database.NewDB(configConfig)
	validate := validator.New()
	authenticationServiceImpl := authentication_service.NewAuthenticationServiceImpl(accountRepositoryImpl, userRepositoryImpl, db, validate)
	authenticationControllerImpl := authentication_controller.NewAuthenticationControllerImpl(authenticationServiceImpl)
	categoryRepositoryImpl := category_repository.NewCategoryRepositoryImpl()
	courseCategoryServiceImpl := course_service.NewCourseCategoryServiceImpl(categoryRepositoryImpl, db)
	courseCategoryControllerImpl := course_controller.NewCourseCategoryControllerImpl(courseCategoryServiceImpl)
	courseRepositoryImpl := course_repository.NewCourseRepositoryImpl()
	coursePopularServiceImpl := course_service.NewCoursePopularServiceImpl(courseRepositoryImpl, db)
	popularCourseControllerImpl := course_controller.NewPopularCourseControllerImpl(coursePopularServiceImpl)
	courseDetailServiceImpl := course_service.NewCourseDetailServiceImpl(courseRepositoryImpl, db)
	detailCourseControllerImpl := course_controller.NewDetailCourseControllerImpl(courseDetailServiceImpl)
	courseReviewServiceImpl := course_service.NewCourseReviewServiceImpl(courseRepositoryImpl, db)
	courseReviewControllerImpl := course_controller.NewCourseReviewControllerImpl(courseReviewServiceImpl)
	webinarSessionRepositoryImpl := webinar_session_repository.NewWebinarSessionRepositoryImpl()
	webinarSessionServiceImpl := webinar_session_service.NewWebinarSessionServiceImpl(webinarSessionRepositoryImpl, db)
	webinarSessionControllerImpl := webinar_session_controller.NewWebinarSessionControllerImpl(webinarSessionServiceImpl)
	elearningModuleRepositoryImpl := elearning_module_repository.NewElearningModuleRepositoryImpl()
	elearningModuleServiceImpl := elearning_module_service.NewElearningModuleServiceImpl(elearningModuleRepositoryImpl, validate, db)
	elearningModuleControllerImpl := elearning_module_controller.NewElearningModuleControllerImpl(elearningModuleServiceImpl)
	quizRepositoryImpl := quiz_repository.NewQuizRepositoryImpl()
	quizServiceImpl := quiz_service.NewQuizServiceImpl(quizRepositoryImpl, validate, db)
	quizControllerImpl := quiz_controller.NewQuizControllerImpl(quizServiceImpl)
	comingSoonCourseServiceImpl := course_service.NewComingSoonCourseServiceImpl(courseRepositoryImpl, db)
	comingSoonCourseControllerImpl := course_controller.NewComingSoonCourseControllerImpl(comingSoonCourseServiceImpl)
	courseSummaryServiceImpl := course_service.NewCourseSummaryServiceImpl(courseRepositoryImpl, db)
	courseSummaryControllerImpl := course_controller.NewCourseSummaryControllerImpl(courseSummaryServiceImpl)
	orderCourseServiceImpl := course_service.NewOrderCourseServiceImpl(courseRepositoryImpl, db)
	orderCourseControllerImpl := course_controller.NewOrderCourseControllerImpl(orderCourseServiceImpl)
	industryInsightRepositoryImpl := industry_insight_repository.NewIndustryInsightRepositoryImpl()
	industryInsightServiceImpl := industry_insight_service.NewIndustryInsightServiceImpl(industryInsightRepositoryImpl, db)
	industryInsightControllerImpl := industry_insight_controller.NewIndustryInsightControllerImpl(industryInsightServiceImpl)
	expertRepositoryImpl := expert_repository.NewExpertRepositoryImpl()
	expertServiceImpl := expert_service.NewExpertServiceImpl(expertRepositoryImpl, db)
	expertControllerImpl := expert_controller.NewExpertControllerImpl(expertServiceImpl)
	userServiceImpl := user_service.NewUserServiceImpl(userRepositoryImpl, db)
	userControllerImpl := user_controller.NewUserControllerImpl(userServiceImpl)
	qnaRepositoryImpl := qna_repository.NewQnaRepositoryImpl()
	qnaServiceImpl := qna_service.NewQnaServiceImpl(qnaRepositoryImpl, validate, db)
	qnaControllerImpl := qna_controller.NewQnaControllerImpl(qnaServiceImpl)
	cartRepositoryImpl := cart_repository.NewCartRepositoryImpl()
	cartServiceImpl := cart_service.NewCartServiceImpl(cartRepositoryImpl, validate, db)
	cartControllerImpl := cart_controller.NewCartControllerImpl(cartServiceImpl)
	echoEcho := app.InitServerWithEcho(authenticationControllerImpl, courseCategoryControllerImpl, popularCourseControllerImpl, detailCourseControllerImpl, courseReviewControllerImpl, webinarSessionControllerImpl, elearningModuleControllerImpl, quizControllerImpl, comingSoonCourseControllerImpl, courseSummaryControllerImpl, orderCourseControllerImpl, industryInsightControllerImpl, expertControllerImpl, userControllerImpl, qnaControllerImpl, cartControllerImpl)
	return echoEcho
}

func InitializedEchoServerForTest() *echo.Echo {
	accountRepositoryImpl := account_repository.NewAccountRepositoryImpl()
	userRepositoryImpl := user_repository.NewUserRepositoryImpl()
	db := database.NewTestDB()
	validate := validator.New()
	authenticationServiceImpl := authentication_service.NewAuthenticationServiceImpl(accountRepositoryImpl, userRepositoryImpl, db, validate)
	authenticationControllerImpl := authentication_controller.NewAuthenticationControllerImpl(authenticationServiceImpl)
	categoryRepositoryImpl := category_repository.NewCategoryRepositoryImpl()
	courseCategoryServiceImpl := course_service.NewCourseCategoryServiceImpl(categoryRepositoryImpl, db)
	courseCategoryControllerImpl := course_controller.NewCourseCategoryControllerImpl(courseCategoryServiceImpl)
	courseRepositoryImpl := course_repository.NewCourseRepositoryImpl()
	coursePopularServiceImpl := course_service.NewCoursePopularServiceImpl(courseRepositoryImpl, db)
	popularCourseControllerImpl := course_controller.NewPopularCourseControllerImpl(coursePopularServiceImpl)
	courseDetailServiceImpl := course_service.NewCourseDetailServiceImpl(courseRepositoryImpl, db)
	detailCourseControllerImpl := course_controller.NewDetailCourseControllerImpl(courseDetailServiceImpl)
	courseReviewServiceImpl := course_service.NewCourseReviewServiceImpl(courseRepositoryImpl, db)
	courseReviewControllerImpl := course_controller.NewCourseReviewControllerImpl(courseReviewServiceImpl)
	webinarSessionRepositoryImpl := webinar_session_repository.NewWebinarSessionRepositoryImpl()
	webinarSessionServiceImpl := webinar_session_service.NewWebinarSessionServiceImpl(webinarSessionRepositoryImpl, db)
	webinarSessionControllerImpl := webinar_session_controller.NewWebinarSessionControllerImpl(webinarSessionServiceImpl)
	elearningModuleRepositoryImpl := elearning_module_repository.NewElearningModuleRepositoryImpl()
	elearningModuleServiceImpl := elearning_module_service.NewElearningModuleServiceImpl(elearningModuleRepositoryImpl, validate, db)
	elearningModuleControllerImpl := elearning_module_controller.NewElearningModuleControllerImpl(elearningModuleServiceImpl)
	quizRepositoryImpl := quiz_repository.NewQuizRepositoryImpl()
	quizServiceImpl := quiz_service.NewQuizServiceImpl(quizRepositoryImpl, validate, db)
	quizControllerImpl := quiz_controller.NewQuizControllerImpl(quizServiceImpl)
	comingSoonCourseServiceImpl := course_service.NewComingSoonCourseServiceImpl(courseRepositoryImpl, db)
	comingSoonCourseControllerImpl := course_controller.NewComingSoonCourseControllerImpl(comingSoonCourseServiceImpl)
	courseSummaryServiceImpl := course_service.NewCourseSummaryServiceImpl(courseRepositoryImpl, db)
	courseSummaryControllerImpl := course_controller.NewCourseSummaryControllerImpl(courseSummaryServiceImpl)
	orderCourseServiceImpl := course_service.NewOrderCourseServiceImpl(courseRepositoryImpl, db)
	orderCourseControllerImpl := course_controller.NewOrderCourseControllerImpl(orderCourseServiceImpl)
	industryInsightRepositoryImpl := industry_insight_repository.NewIndustryInsightRepositoryImpl()
	industryInsightServiceImpl := industry_insight_service.NewIndustryInsightServiceImpl(industryInsightRepositoryImpl, db)
	industryInsightControllerImpl := industry_insight_controller.NewIndustryInsightControllerImpl(industryInsightServiceImpl)
	expertRepositoryImpl := expert_repository.NewExpertRepositoryImpl()
	expertServiceImpl := expert_service.NewExpertServiceImpl(expertRepositoryImpl, db)
	expertControllerImpl := expert_controller.NewExpertControllerImpl(expertServiceImpl)
	userServiceImpl := user_service.NewUserServiceImpl(userRepositoryImpl, db)
	userControllerImpl := user_controller.NewUserControllerImpl(userServiceImpl)
	qnaRepositoryImpl := qna_repository.NewQnaRepositoryImpl()
	qnaServiceImpl := qna_service.NewQnaServiceImpl(qnaRepositoryImpl, validate, db)
	qnaControllerImpl := qna_controller.NewQnaControllerImpl(qnaServiceImpl)
	cartRepositoryImpl := cart_repository.NewCartRepositoryImpl()
	cartServiceImpl := cart_service.NewCartServiceImpl(cartRepositoryImpl, validate, db)
	cartControllerImpl := cart_controller.NewCartControllerImpl(cartServiceImpl)
	echoEcho := app.InitServerWithEcho(authenticationControllerImpl, courseCategoryControllerImpl, popularCourseControllerImpl, detailCourseControllerImpl, courseReviewControllerImpl, webinarSessionControllerImpl, elearningModuleControllerImpl, quizControllerImpl, comingSoonCourseControllerImpl, courseSummaryControllerImpl, orderCourseControllerImpl, industryInsightControllerImpl, expertControllerImpl, userControllerImpl, qnaControllerImpl, cartControllerImpl)
	return echoEcho
}

// injector.go:

var userRepositorySet = wire.NewSet(user_repository.NewUserRepositoryImpl, wire.Bind(new(user_repository.UserRepository), new(*user_repository.UserRepositoryImpl)))

var accountRepositorySet = wire.NewSet(account_repository.NewAccountRepositoryImpl, wire.Bind(new(account_repository.AccountRepository), new(*account_repository.AccountRepositoryImpl)))

var authenticationSet = wire.NewSet(authentication_service.NewAuthenticationServiceImpl, wire.Bind(new(authentication_service.AuthenticationService), new(*authentication_service.AuthenticationServiceImpl)), authentication_controller.NewAuthenticationControllerImpl, wire.Bind(new(authentication_controller.AuthenticationController), new(*authentication_controller.AuthenticationControllerImpl)))

var courseCategorySet = wire.NewSet(category_repository.NewCategoryRepositoryImpl, wire.Bind(new(category_repository.CategoryRepository), new(*category_repository.CategoryRepositoryImpl)), course_service.NewCourseCategoryServiceImpl, wire.Bind(new(course_service.CourseCategoryService), new(*course_service.CourseCategoryServiceImpl)), course_controller.NewCourseCategoryControllerImpl, wire.Bind(new(course_controller.CourseCategoryController), new(*course_controller.CourseCategoryControllerImpl)))

var popularCourseSet = wire.NewSet(course_service.NewCoursePopularServiceImpl, wire.Bind(new(course_service.CoursePopularService), new(*course_service.CoursePopularServiceImpl)), course_controller.NewPopularCourseControllerImpl, wire.Bind(new(course_controller.PopularCourseController), new(*course_controller.PopularCourseControllerImpl)))

var courseRepositorySet = wire.NewSet(course_repository.NewCourseRepositoryImpl, wire.Bind(new(course_repository.CourseRepository), new(*course_repository.CourseRepositoryImpl)))

var detailCourseSet = wire.NewSet(course_controller.NewDetailCourseControllerImpl, wire.Bind(new(course_controller.DetailCourseController), new(*course_controller.DetailCourseControllerImpl)), course_service.NewCourseDetailServiceImpl, wire.Bind(new(course_service.CourseDetailService), new(*course_service.CourseDetailServiceImpl)))

var reviewCourseSet = wire.NewSet(course_controller.NewCourseReviewControllerImpl, wire.Bind(new(course_controller.CourseReviewController), new(*course_controller.CourseReviewControllerImpl)), course_service.NewCourseReviewServiceImpl, wire.Bind(new(course_service.CourseReviewService), new(*course_service.CourseReviewServiceImpl)))

var webinarSessionSet = wire.NewSet(webinar_session_repository.NewWebinarSessionRepositoryImpl, wire.Bind(new(webinar_session_repository.WebinarSessionRepository), new(*webinar_session_repository.WebinarSessionRepositoryImpl)), webinar_session_service.NewWebinarSessionServiceImpl, wire.Bind(new(webinar_session_service.WebinarSessionService), new(*webinar_session_service.WebinarSessionServiceImpl)), webinar_session_controller.NewWebinarSessionControllerImpl, wire.Bind(new(webinar_session_controller.WebinarSessionController), new(*webinar_session_controller.WebinarSessionControllerImpl)))

var elearningModuleSet = wire.NewSet(elearning_module_repository.NewElearningModuleRepositoryImpl, wire.Bind(new(elearning_module_repository.ElearningModuleRepository), new(*elearning_module_repository.ElearningModuleRepositoryImpl)), elearning_module_service.NewElearningModuleServiceImpl, wire.Bind(new(elearning_module_service.ElearningModuleService), new(*elearning_module_service.ElearningModuleServiceImpl)), elearning_module_controller.NewElearningModuleControllerImpl, wire.Bind(new(elearning_module_controller.ElearningModuleController), new(*elearning_module_controller.ElearningModuleControllerImpl)))

var quizSet = wire.NewSet(quiz_repository.NewQuizRepositoryImpl, wire.Bind(new(quiz_repository.QuizRepository), new(*quiz_repository.QuizRepositoryImpl)), quiz_service.NewQuizServiceImpl, wire.Bind(new(quiz_service.QuizService), new(*quiz_service.QuizServiceImpl)), quiz_controller.NewQuizControllerImpl, wire.Bind(new(quiz_controller.QuizController), new(*quiz_controller.QuizControllerImpl)))

var comingSoonCourseSet = wire.NewSet(course_service.NewComingSoonCourseServiceImpl, wire.Bind(new(course_service.ComingSoonCourseService), new(*course_service.ComingSoonCourseServiceImpl)), course_controller.NewComingSoonCourseControllerImpl, wire.Bind(new(course_controller.ComingSoonCourseController), new(*course_controller.ComingSoonCourseControllerImpl)))

var courseSummarySet = wire.NewSet(course_service.NewCourseSummaryServiceImpl, wire.Bind(new(course_service.CourseSummaryService), new(*course_service.CourseSummaryServiceImpl)), course_controller.NewCourseSummaryControllerImpl, wire.Bind(new(course_controller.CourseSummaryController), new(*course_controller.CourseSummaryControllerImpl)))

var orderCourseSet = wire.NewSet(course_service.NewOrderCourseServiceImpl, wire.Bind(new(course_service.OrderCourseService), new(*course_service.OrderCourseServiceImpl)), course_controller.NewOrderCourseControllerImpl, wire.Bind(new(course_controller.OrderCourseController), new(*course_controller.OrderCourseControllerImpl)))

var industryInsightSet = wire.NewSet(industry_insight_repository.NewIndustryInsightRepositoryImpl, wire.Bind(new(industry_insight_repository.IndustryInsightRepository), new(*industry_insight_repository.IndustryInsightRepositoryImpl)), industry_insight_service.NewIndustryInsightServiceImpl, wire.Bind(new(industry_insight_service.IndustryInsightService), new(*industry_insight_service.IndustryInsightServiceImpl)), industry_insight_controller.NewIndustryInsightControllerImpl, wire.Bind(new(industry_insight_controller.IndustryInsightController), new(*industry_insight_controller.IndustryInsightControllerImpl)))

var expertSet = wire.NewSet(expert_repository.NewExpertRepositoryImpl, wire.Bind(new(expert_repository.ExpertRepository), new(*expert_repository.ExpertRepositoryImpl)), expert_service.NewExpertServiceImpl, wire.Bind(new(expert_service.ExpertService), new(*expert_service.ExpertServiceImpl)), expert_controller.NewExpertControllerImpl, wire.Bind(new(expert_controller.ExpertController), new(*expert_controller.ExpertControllerImpl)))

var userSet = wire.NewSet(user_service.NewUserServiceImpl, wire.Bind(new(user_service.UserService), new(*user_service.UserServiceImpl)), user_controller.NewUserControllerImpl, wire.Bind(new(user_controller.UserController), new(*user_controller.UserControllerImpl)))

var qnaSet = wire.NewSet(qna_repository.NewQnaRepositoryImpl, wire.Bind(new(qna_repository.QnaRepository), new(*qna_repository.QnaRepositoryImpl)), qna_service.NewQnaServiceImpl, wire.Bind(new(qna_service.QnaService), new(*qna_service.QnaServiceImpl)), qna_controller.NewQnaControllerImpl, wire.Bind(new(qna_controller.QnaController), new(*qna_controller.QnaControllerImpl)))

var cartSet = wire.NewSet(cart_repository.NewCartRepositoryImpl, wire.Bind(new(cart_repository.CartRepository), new(*cart_repository.CartRepositoryImpl)), cart_service.NewCartServiceImpl, wire.Bind(new(cart_service.CartService), new(*cart_service.CartServiceImpl)), cart_controller.NewCartControllerImpl, wire.Bind(new(cart_controller.CartController), new(*cart_controller.CartControllerImpl)))

var completeSet = wire.NewSet(app.InitServerWithEcho, validator.New, config.LoadConfigFromEnv, accountRepositorySet,
	userRepositorySet,
	authenticationSet,
	courseRepositorySet,
	courseCategorySet,
	popularCourseSet,
	detailCourseSet,
	reviewCourseSet,
	webinarSessionSet,
	elearningModuleSet,
	quizSet,
	comingSoonCourseSet,
	courseSummarySet,
	orderCourseSet,
	industryInsightSet,
	expertSet,
	userSet,
	qnaSet,
	cartSet,
)
