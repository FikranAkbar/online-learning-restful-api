package cart_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/payment"
	"online-learning-restful-api/service/payment_service"
)

type PaymentControllerImpl struct {
	payment_service.PaymentService
}

func NewPaymentControllerImpl(paymentService payment_service.PaymentService) *PaymentControllerImpl {
	return &PaymentControllerImpl{PaymentService: paymentService}
}

func (controller *PaymentControllerImpl) CreateNewCourseOrder(c echo.Context) error {
	var courseOrderRequest payment.CourseOrderRequest
	err := c.Bind(&courseOrderRequest)
	helper.PanicIfError(err)

	courseOrderResponse := controller.PaymentService.CreateNewCourseOrder(c.Request().Context(), courseOrderRequest.CourseIds)

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    courseOrderResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
