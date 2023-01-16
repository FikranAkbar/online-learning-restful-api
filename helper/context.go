package helper

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func ExtractUserInfoFromJwtToContext(c echo.Context) context.Context {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		panic("JWT token missing or invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		panic("Failed cast claims to jwt.MapClaims")
	}

	ctx := context.WithValue(c.Request().Context(), "userInfo", claims)

	return ctx
}
