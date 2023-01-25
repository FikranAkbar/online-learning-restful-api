package qna_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/qna"
	"online-learning-restful-api/service/qna_service"
	"strconv"
)

type QnaControllerImpl struct {
	qna_service.QnaService
}

func NewQnaControllerImpl(qnaService qna_service.QnaService) *QnaControllerImpl {
	return &QnaControllerImpl{QnaService: qnaService}
}

func (controller *QnaControllerImpl) GetQnaQuestionsByCourseId(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	qnaQuestionsResponse := controller.QnaService.GetQnaQuestionsByCourseId(c.Request().Context(), uint(courseId))

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    qnaQuestionsResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *QnaControllerImpl) CreateNewQnaQuestion(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	var qnaQuestionRequest qna.CreateQnaQuestionRequest
	err = c.Bind(&qnaQuestionRequest)
	helper.PanicIfError(err)

	shortQnaQuestionResponse := controller.QnaService.CreateNewQnaQuestion(c.Request().Context(), uint(courseId), qnaQuestionRequest)

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    shortQnaQuestionResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *QnaControllerImpl) GetDetailQnaQuestionByQnaQuestionId(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	qnaQuestionId, err := strconv.Atoi(c.Param("qnaQuestionId"))
	helper.PanicIfError(err)

	detailQnaQuestionResponse := controller.QnaService.GetDetailQnaQuestionByQnaQuestionId(
		c.Request().Context(), uint(courseId), uint(qnaQuestionId))

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    detailQnaQuestionResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *QnaControllerImpl) CreateNewQnaAnswer(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	qnaQuestionId, err := strconv.Atoi(c.Param("qnaQuestionId"))
	helper.PanicIfError(err)

	var request qna.CreateQnaAnswerRequest
	err = c.Bind(&request)
	helper.PanicIfError(err)

	shortQnaAnswerResponse := controller.QnaService.CreateNewQnaAnswer(c.Request().Context(), uint(courseId), uint(qnaQuestionId), request)

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    shortQnaAnswerResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

