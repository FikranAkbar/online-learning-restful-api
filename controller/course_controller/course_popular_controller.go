package course_controller

import "github.com/labstack/echo/v4"

type CoursePopularController interface {
	GetPopularCourses(c echo.Context) error
}
