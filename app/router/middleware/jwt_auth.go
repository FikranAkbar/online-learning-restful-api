package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/helper"
	"strings"
)

func JWTAuthorization(echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			return exception.GenerateHTTPError(exception.BadRequest, "Invalid token")
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != helper.JwtSigningMethod {
				return nil, fmt.Errorf("signing method invalid")
			}

			return helper.JwtSignatureKey, nil
		})

		if err != nil {
			return exception.GenerateHTTPError(exception.BadRequest, err.Error())
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return exception.GenerateHTTPError(exception.BadRequest, "Parsing to jwt map claims failed")
		}

		ctx := context.WithValue(context.Background(), "userInfo", claims)
		c.Request().WithContext(ctx)

		return nil
	}
}
