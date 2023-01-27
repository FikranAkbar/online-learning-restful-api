package payment_controller

import "github.com/labstack/echo/v4"

type PaymentController interface {
	CreateNewCourseOrder(c echo.Context) error
}
