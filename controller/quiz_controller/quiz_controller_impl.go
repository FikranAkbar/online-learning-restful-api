package quiz_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/quiz"
	"online-learning-restful-api/service/quiz_service"
	"strconv"
)

type QuizControllerImpl struct {
	quiz_service.QuizService
}

func NewQuizControllerImpl(quizService quiz_service.QuizService) *QuizControllerImpl {
	return &QuizControllerImpl{QuizService: quizService}
}

func (controller QuizControllerImpl) GetQuizAnswersByModuleId(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	elearningModuleId, err := strconv.Atoi(c.Param("moduleId"))
	helper.PanicIfError(err)

	quizAnswersResponse := controller.QuizService.GetQuizAnswersByModuleId(c.Request().Context(), uint(courseId), uint(elearningModuleId))

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    quizAnswersResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller QuizControllerImpl) CreateNewQuizAnswer(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	elearningModuleId, err := strconv.Atoi(c.Param("moduleId"))
	helper.PanicIfError(err)

	var request quiz.ShortQuizAnswerRequest
	err = c.Bind(&request)
	helper.PanicIfError(err)

	quizAnswerResponse := controller.QuizService.CreateNewQuizAnswer(c.Request().Context(), uint(courseId), uint(elearningModuleId), request)

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    quizAnswerResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
