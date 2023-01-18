package course_controller

import "github.com/labstack/echo/v4"

type CourseCategoryController interface {
	GetAllCourseCategories(c echo.Context) error
	GetCoursesByCategoryId(c echo.Context) error
}
