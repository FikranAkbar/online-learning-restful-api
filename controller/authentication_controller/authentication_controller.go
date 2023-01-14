package authentication_controller

import "github.com/labstack/echo/v4"

type AuthenticationController interface {
	LoginUserWithEmailPassword(c echo.Context) error
	RegisterUserWithEmailPassword(c echo.Context) error
}
