package industry_insight_controller

import "github.com/labstack/echo/v4"

type IndustryInsightController interface {
	GetIndustryInsights(c echo.Context) error
	GetIndustryInsightById(c echo.Context) error
}
