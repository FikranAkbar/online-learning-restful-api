package course_controller

import "github.com/labstack/echo/v4"

type CourseReviewController interface {
	GetCourseReviewsByCourseId(c echo.Echo) error
}
