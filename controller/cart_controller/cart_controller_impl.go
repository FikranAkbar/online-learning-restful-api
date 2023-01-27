package cart_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/payment"
	"online-learning-restful-api/service/cart_service"
)

type CartControllerImpl struct {
	cart_service.CartService
}

func NewCartControllerImpl(paymentService cart_service.CartService) *CartControllerImpl {
	return &CartControllerImpl{CartService: paymentService}
}

func (controller *CartControllerImpl) BuyCartItems(c echo.Context) error {
	var courseOrderRequest payment.CourseOrderRequest
	err := c.Bind(&courseOrderRequest)
	helper.PanicIfError(err)

	courseOrderResponse := controller.CartService.BuyCartItems(c.Request().Context(), courseOrderRequest.CourseIds)

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    courseOrderResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
