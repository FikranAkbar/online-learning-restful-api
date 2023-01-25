package qna_controller

import "github.com/labstack/echo/v4"

type QnaController interface {
	GetQnaQuestionsByCourseId(c echo.Context) error
	CreateNewQnaQuestion(c echo.Context) error
	GetDetailQnaQuestionByQnaQuestionId(c echo.Context) error
	CreateNewQnaAnswer(c echo.Context) error
}
