package exception

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/model/web"
	"strings"
)

const (
	Unauthorized  = "unauthorized"
	NotFoundError = "not found"
)

func ErrorHandler(err error, c echo.Context) {
	if strings.Contains(strings.ToLower(err.Error()), NotFoundError) {
		recordNotFoundError(err, c)
		return
	}

	if strings.Contains(strings.ToLower(err.Error()), Unauthorized) {
		unauthorizedError(err, c)
		return
	}

	internalServerError(err, c)
}

func internalServerError(err error, c echo.Context) {
	response := web.APIResponse{
		Status:  500,
		Message: "INTERNAL SERVER ERROR",
		Data:    err.Error(),
	}

	_ = c.JSON(http.StatusInternalServerError, response)
}

func recordNotFoundError(err error, c echo.Context) {
	response := web.APIResponse{
		Status:  404,
		Message: "NOT FOUND",
		Data:    err.Error(),
	}

	_ = c.JSON(http.StatusInternalServerError, response)
}

func unauthorizedError(err error, c echo.Context) {
	response := web.APIResponse{
		Status:  401,
		Message: "UNAUTHORIZED",
		Data:    err.Error(),
	}

	_ = c.JSON(http.StatusUnauthorized, response)
}
