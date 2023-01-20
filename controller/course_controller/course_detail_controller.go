package course_controller

import "github.com/labstack/echo/v4"

type CourseDetailController interface {
	GetCoursesByKeyword(c echo.Context) error
	GetDetailCourseByCourseId(c echo.Context) error
	GetUserCourseProgressionByCourseId(c echo.Context) error
}
