package industry_insight_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/service/industry_insight_service"
	"strconv"
)

type IndustryInsightControllerImpl struct {
	industry_insight_service.IndustryInsightService
}

func NewIndustryInsightControllerImpl(industryInsightService industry_insight_service.IndustryInsightService) *IndustryInsightControllerImpl {
	return &IndustryInsightControllerImpl{IndustryInsightService: industryInsightService}
}

func (controller *IndustryInsightControllerImpl) GetIndustryInsights(c echo.Context) error {
	industryInsightsResponse := controller.IndustryInsightService.GetIndustryInsights(c.Request().Context())

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    industryInsightsResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *IndustryInsightControllerImpl) GetIndustryInsightById(c echo.Context) error {
	industryInsightId, err := strconv.Atoi(c.Param("industryInsightId"))
	helper.PanicIfError(err)

	industryInsightResponse := controller.IndustryInsightService.GetIndustryInsightById(c.Request().Context(), uint(industryInsightId))

	apiResponse := web.APIResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    industryInsightResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
