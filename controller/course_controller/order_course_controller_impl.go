package course_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/service/course_service"
	"strconv"
)

type OrderCourseControllerImpl struct {
	course_service.OrderCourseService
}

func NewOrderCourseControllerImpl(orderCourseService course_service.OrderCourseService) *OrderCourseControllerImpl {
	return &OrderCourseControllerImpl{OrderCourseService: orderCourseService}
}

func (controller *OrderCourseControllerImpl) CreateNewCourseOrder(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	orderCourseResponse := controller.OrderCourseService.CreateNewCourseOrder(c.Request().Context(), uint(courseId))

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    orderCourseResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *OrderCourseControllerImpl) DeleteCourseOrder(c echo.Context) error {
	courseId, err := strconv.Atoi(c.Param("courseId"))
	helper.PanicIfError(err)

	orderCourseResponse := controller.OrderCourseService.DeleteCourseOrder(c.Request().Context(), uint(courseId))

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    orderCourseResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

