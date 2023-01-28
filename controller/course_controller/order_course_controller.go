package course_controller

import "github.com/labstack/echo/v4"

type OrderCourseController interface {
	CreateNewCourseOrder(c echo.Context) error
	DeleteCourseOrder(c echo.Context) error
}
