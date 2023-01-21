package elearning_module_controller

import "github.com/labstack/echo/v4"

type ElearningModuleController interface {
	GetOverviewElearningModulesByCourseId(c echo.Context) error
}
