package quiz_controller

import "github.com/labstack/echo/v4"

type QuizController interface {
	GetQuizAnswersByModuleId(c echo.Context) error
	CreateNewQuizAnswer(c echo.Context) error
}
