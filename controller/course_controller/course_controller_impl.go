package course_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/service/course_service"
	"strconv"
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

func (controller *CourseControllerImpl) GetDetailCourseByCourseId(c echo.Context) error {
	courseId, courseIdErr := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(courseIdErr)

	var detailCourseResponse course.DetailCourseResponse

	userIdQuery := c.QueryParam("userid")
	if userIdQuery == "" {
		detailCourseResponse = controller.CourseService.GetDetailCourseByCourseId(c.Request().Context(), uint(courseId), nil)
	} else {
		userId, userIdErr := strconv.Atoi(userIdQuery)
		helper.PanicIfError(userIdErr)
		newUserId := uint(userId)
		detailCourseResponse = controller.CourseService.GetDetailCourseByCourseId(c.Request().Context(), uint(courseId), &newUserId)
	}

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    detailCourseResponse,
	}

	return c.JSON(200, apiResponse)
}
