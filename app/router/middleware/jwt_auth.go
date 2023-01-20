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

type UserTokenInfo struct {
	UserId    uint
	UserEmail string
	UserName  string
}

var (
	ContextUserInfoKey    = "user_info"
	UnauthorizedErrorInfo = exception.Unauthorized + ". User token was not found"
)

func JWTAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			return exception.GenerateHTTPError(exception.Unauthorized, "Invalid token")
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
			return exception.GenerateHTTPError(exception.Unauthorized, err.Error())
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return exception.GenerateHTTPError(exception.Unauthorized, "Parsing to jwt map claims failed")
		}

		userId := uint(claims["user_id"].(float64))
		userEmail := claims["user_email"].(string)
		userName := claims["user_name"].(string)

		userTokenInfo := UserTokenInfo{
			UserId:    userId,
			UserEmail: userEmail,
			UserName:  userName,
		}

		ctx := context.WithValue(c.Request().Context(), ContextUserInfoKey, userTokenInfo)
		newRequest := c.Request().WithContext(ctx)
		c.SetRequest(newRequest)

		return next(c)
	}
}
