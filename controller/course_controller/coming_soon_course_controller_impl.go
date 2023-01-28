package course_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/service/course_service"
)

type ComingSoonCourseControllerImpl struct {
	course_service.ComingSoonCourseService
}

func NewComingSoonCourseControllerImpl(comingSoonCourseService course_service.ComingSoonCourseService) *ComingSoonCourseControllerImpl {
	return &ComingSoonCourseControllerImpl{ComingSoonCourseService: comingSoonCourseService}
}

func (controller *ComingSoonCourseControllerImpl) GetComingSoonCourses(c echo.Context) error {
	comingSoonCoursesResponse := controller.ComingSoonCourseService.GetComingSoonCourses(c.Request().Context())

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    comingSoonCoursesResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
