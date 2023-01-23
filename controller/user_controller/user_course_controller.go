package user_controller

import "github.com/labstack/echo/v4"

type UserController interface {
	GetUserCourses(c echo.Context) error
}
