package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitMiddleware(e *echo.Echo) {
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowMethods: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	//	AllowHeaders: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	//}))
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
}
