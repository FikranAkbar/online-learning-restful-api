package elearning_module_controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/video"
	"online-learning-restful-api/service/elearning_module_service"
	"strconv"
)

type ElearningModuleControllerImpl struct {
	elearning_module_service.ElearningModuleService
}

func NewElearningModuleControllerImpl(service elearning_module_service.ElearningModuleService) *ElearningModuleControllerImpl {
	return &ElearningModuleControllerImpl{ElearningModuleService: service}
}

func (controller *ElearningModuleControllerImpl) GetOverviewElearningModulesByCourseId(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	overviewElearningModulesResponse := controller.ElearningModuleService.
		GetOverviewElearningModulesByCourseId(c.Request().Context(), uint(courseId))

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    overviewElearningModulesResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *ElearningModuleControllerImpl) GetDetailElearningModuleByElearningModuleId(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	elearningModuleId, err := strconv.Atoi(c.Param("moduleId"))
	helper.PanicIfError(err)

	detailElearningModuleResponse := controller.ElearningModuleService.
		GetDetailElearningModuleByElearningModuleId(c.Request().Context(), uint(courseId), uint(elearningModuleId))

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    detailElearningModuleResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *ElearningModuleControllerImpl) SaveVideoProgressionInModule(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	elearningModuleId, err := strconv.Atoi(c.Param("moduleId"))
	helper.PanicIfError(err)

	var request video.UserVideoProgressionRequest
	err = c.Bind(&request)
	helper.PanicIfError(err)

	fmt.Println("Request: Progression", request.Progression, "Is Complete", request.IsComplete)

	userVideoProgressionResponse := controller.ElearningModuleService.SaveVideoProgressionInModule(
		c.Request().Context(), uint(courseId), uint(elearningModuleId), request)

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    userVideoProgressionResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
