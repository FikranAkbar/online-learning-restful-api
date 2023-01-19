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
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)
	newCourseId := uint(courseId)

	userId, err := strconv.Atoi(c.QueryParam("userid"))

	var detailCourseResponse course.DetailCourseResponse

	if err != nil {
		newUserId := uint(userId)
		detailCourseResponse = controller.CourseService.GetDetailCourseByCourseId(c.Request().Context(), newCourseId, &newUserId)
	} else {
		detailCourseResponse = controller.CourseService.GetDetailCourseByCourseId(c.Request().Context(), newCourseId, nil)

	}

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    detailCourseResponse,
	}

	return c.JSON(200, apiResponse)
}
