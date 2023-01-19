package course_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/service/course_service"
)

type CoursePopularControllerImpl struct {
	course_service.CoursePopularService
}

func NewCoursePopularControllerImpl(service course_service.CoursePopularService) *CoursePopularControllerImpl {
	return &CoursePopularControllerImpl{service}
}

func (controller *CoursePopularControllerImpl) GetPopularCourses(c echo.Context) error {
	coursesResponse := controller.CoursePopularService.GetPopularCourses(c.Request().Context())

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    coursesResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

