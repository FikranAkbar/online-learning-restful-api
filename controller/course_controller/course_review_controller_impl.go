package course_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/service/course_service"
	"strconv"
)

type CourseReviewControllerImpl struct {
	course_service.CourseReviewService
}

func NewCourseReviewControllerImpl(service course_service.CourseReviewService) *CourseReviewControllerImpl {
	return &CourseReviewControllerImpl{CourseReviewService: service}
}

func (controller *CourseReviewControllerImpl) GetCourseReviewsByCourseId(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	reviewCourseResponses := controller.CourseReviewService.GetCourseReviewsByCourseId(c.Request().Context(), uint(courseId))

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    reviewCourseResponses,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
