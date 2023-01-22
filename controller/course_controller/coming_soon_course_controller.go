package course_controller

import "github.com/labstack/echo/v4"

type ComingSoonCourseController interface {
	GetComingSoonCourses(c echo.Context) error
}
