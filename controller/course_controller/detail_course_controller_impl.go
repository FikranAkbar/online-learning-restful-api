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

type DetailCourseControllerImpl struct {
	course_service.CourseDetailService
}

func NewDetailCourseControllerImpl(courseService course_service.CourseDetailService) *DetailCourseControllerImpl {
	return &DetailCourseControllerImpl{CourseDetailService: courseService}
}

func (controller *DetailCourseControllerImpl) GetCoursesByKeyword(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	coursesResponse := controller.CourseDetailService.GetCoursesByKeyword(c.Request().Context(), keyword)

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    coursesResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *DetailCourseControllerImpl) GetDetailCourseByCourseId(c echo.Context) error {
	courseId, courseIdErr := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(courseIdErr)

	var detailCourseResponse course.DetailCourseResponse

	userIdQuery := c.QueryParam("userid")
	if userIdQuery == "" {
		detailCourseResponse = controller.CourseDetailService.GetDetailCourseByCourseId(c.Request().Context(), uint(courseId), nil)
	} else {
		userId, userIdErr := strconv.Atoi(userIdQuery)
		helper.PanicIfError(userIdErr)
		newUserId := uint(userId)
		detailCourseResponse = controller.CourseDetailService.GetDetailCourseByCourseId(c.Request().Context(), uint(courseId), &newUserId)
	}

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    detailCourseResponse,
	}

	return c.JSON(200, apiResponse)
}

func (controller *DetailCourseControllerImpl) GetUserCourseProgressionByCourseId(c echo.Context) error {
	courseId := c.Param("courseId")
	newCourseId, err := strconv.Atoi(courseId)
	helper.PanicIfError(err)

	userCourseProgressionResponse := controller.CourseDetailService.GetUserCourseProgressionByCourseId(c.Request().Context(), uint(newCourseId))

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    userCourseProgressionResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
