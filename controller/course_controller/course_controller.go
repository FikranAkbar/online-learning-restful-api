package course_controller

import "github.com/labstack/echo/v4"

type CourseController interface {
	GetCoursesByKeyword(c echo.Context) error
	GetDetailCourseByCourseId(c echo.Context) error
}
