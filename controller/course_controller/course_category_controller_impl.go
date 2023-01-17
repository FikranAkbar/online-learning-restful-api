package course_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/service/course_service"
)

type CourseCategoryControllerImpl struct {
	CourseCategoryService course_service.CourseCategoryService
}

func NewCourseCategoryControllerImpl(service course_service.CourseCategoryService) *CourseCategoryControllerImpl {
	return &CourseCategoryControllerImpl{service}
}

func (controller *CourseCategoryControllerImpl) GetAllCourseCategories(c echo.Context) error {
	categoriesResponse := controller.CourseCategoryService.GetAllCategories(c.Request().Context())
	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    categoriesResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
