package course_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/service/course_service"
	"strconv"
)

type CourseSummaryControllerImpl struct {
	course_service.CourseSummaryService
}

func NewCourseSummaryControllerImpl(courseSummaryService course_service.CourseSummaryService) *CourseSummaryControllerImpl {
	return &CourseSummaryControllerImpl{CourseSummaryService: courseSummaryService}
}

func (controller *CourseSummaryControllerImpl) GetCourseSummary(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	courseSummaryResponse := controller.CourseSummaryService.GetCourseSummary(c.Request().Context(), uint(courseId))

	apiResponse := web.APIResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    courseSummaryResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
