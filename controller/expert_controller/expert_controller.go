package expert_controller

import "github.com/labstack/echo/v4"

type ExpertController interface {
	GetExpertDetailById(c echo.Context) error
}
