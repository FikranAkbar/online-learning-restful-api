package course_controller

import "github.com/labstack/echo/v4"

type CourseSummaryController interface {
	GetCourseSummary(c echo.Context) error
}
