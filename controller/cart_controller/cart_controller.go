package cart_controller

import "github.com/labstack/echo/v4"

type CartController interface {
	BuyCartItems(c echo.Context) error
}
