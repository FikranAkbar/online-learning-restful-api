package expert_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/service/expert_service"
	"strconv"
)

type ExpertControllerImpl struct {
	expert_service.ExpertService
}

func NewExpertControllerImpl(expertService expert_service.ExpertService) *ExpertControllerImpl {
	return &ExpertControllerImpl{ExpertService: expertService}
}

func (controller *ExpertControllerImpl) GetExpertDetailById(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("expertId"))
	helper.PanicIfError(err)

	detailExpertResponse := controller.ExpertService.GetExpertDetailById(c.Request().Context(), uint(courseId))

	apiResponse := web.APIResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    detailExpertResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
