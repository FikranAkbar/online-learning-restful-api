package webinar_session_controller

import "github.com/labstack/echo/v4"

type WebinarSessionController interface {
	GetOverviewWebinarSessionsByCourseId(c echo.Context) error
}
