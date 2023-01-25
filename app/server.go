package app

import (
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/app/router"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/controller/authentication_controller"
	"online-learning-restful-api/controller/course_controller"
	"online-learning-restful-api/controller/elearning_module_controller"
	"online-learning-restful-api/controller/expert_controller"
	"online-learning-restful-api/controller/industry_insight_controller"
	"online-learning-restful-api/controller/qna_controller"
	"online-learning-restful-api/controller/quiz_controller"
	"online-learning-restful-api/controller/user_controller"
	"online-learning-restful-api/controller/webinar_session_controller"
	"online-learning-restful-api/exception"
)

func InitServerWithEcho(
	authenticationController authentication_controller.AuthenticationController,
	courseCategoryController course_controller.CourseCategoryController,
	coursePopularController course_controller.PopularCourseController,
	detailCourseController course_controller.DetailCourseController,
	courseReviewController course_controller.CourseReviewController,
	webinarSessionController webinar_session_controller.WebinarSessionController,
	elearningModuleController elearning_module_controller.ElearningModuleController,
	quizController quiz_controller.QuizController,
	comingSoonCourseController course_controller.ComingSoonCourseController,
	courseSummaryController course_controller.CourseSummaryController,
	industryInsightController industry_insight_controller.IndustryInsightController,
	expertController expert_controller.ExpertController,
	userController user_controller.UserController,
	qnaController qna_controller.QnaController,
) *echo.Echo {
	e := echo.New()

	router.InitRoutes(
		authenticationController,
		courseCategoryController,
		coursePopularController,
		detailCourseController,
		courseReviewController,
		webinarSessionController,
		elearningModuleController,
		quizController,
		comingSoonCourseController,
		courseSummaryController,
		industryInsightController,
		expertController,
		userController,
		qnaController,
		e)
	middleware.InitMiddleware(e)

	e.HTTPErrorHandler = exception.ErrorHandler

	return e
}
