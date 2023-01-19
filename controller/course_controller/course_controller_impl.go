package course_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/service/course_service"
)

type CourseControllerImpl struct {
	course_service.CourseService
}

func NewCourseControllerImpl(courseService course_service.CourseService) *CourseControllerImpl {
	return &CourseControllerImpl{CourseService: courseService}
}

func (controller *CourseControllerImpl) GetCoursesByKeyword(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	coursesResponse := controller.CourseService.GetCoursesByKeyword(c.Request().Context(), keyword)

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    coursesResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
