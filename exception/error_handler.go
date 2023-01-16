package exception

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/model/web"
	"strings"
)

const (
	Unauthorized = "Unauthorized"
	NotFound     = "Not found"
	BadRequest   = "Bad request"
	Validation   = "Validation"
	Duplicate    = "Duplicate"
)

func CheckErrorContains(err error, subStr string) bool {
	if err != nil {
		return strings.Contains(strings.ToLower(err.Error()), strings.ToLower(subStr))
	}

	return false
}

func GenerateHTTPError(httpError string, errorLog string) error {
	return fmt.Errorf("%v. %v", httpError, errorLog)
}

func ErrorHandler(err error, c echo.Context) {
	if CheckErrorContains(err, BadRequest) || CheckErrorContains(err, Validation) {
		badRequestError(err, c)
		return
	}

	if CheckErrorContains(err, Unauthorized) {
		unauthorizedError(err, c)
		return
	}

	if CheckErrorContains(err, NotFound) {
		recordNotFoundError(err, c)
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

	_ = c.JSON(http.StatusNotFound, response)
}

func unauthorizedError(err error, c echo.Context) {
	response := web.APIResponse{
		Status:  401,
		Message: "UNAUTHORIZED",
		Data:    err.Error(),
	}

	_ = c.JSON(http.StatusUnauthorized, response)
}

func badRequestError(err error, c echo.Context) {
	response := web.APIResponse{
		Status:  400,
		Message: "BAD REQUEST",
		Data:    err.Error(),
	}

	_ = c.JSON(http.StatusBadRequest, response)
}
