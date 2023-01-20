package webinar_session_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/service/webinar_session_service"
	"strconv"
)

type WebinarSessionControllerImpl struct {
	webinar_session_service.WebinarSessionService
}

func NewWebinarSessionControllerImpl(service webinar_session_service.WebinarSessionService) *WebinarSessionControllerImpl {
	return &WebinarSessionControllerImpl{WebinarSessionService: service}
}

func (controller *WebinarSessionControllerImpl) GetOverviewWebinarSessionsByCourseId(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	overviewWebinarSessionsResponse := controller.WebinarSessionService.
		GetOverviewWebinarSessionsByCourseId(c.Request().Context(), uint(courseId))

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    overviewWebinarSessionsResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

