package authentication_controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/authentication"
	"online-learning-restful-api/service/authentication_service"
)

type AuthenticationControllerImpl struct {
	AuthenticationService authentication_service.AuthenticationService
}

func NewAuthenticationControllerImpl(authenticationService authentication_service.AuthenticationService) *AuthenticationControllerImpl {
	return &AuthenticationControllerImpl{
		AuthenticationService: authenticationService,
	}
}

func (controller *AuthenticationControllerImpl) LoginUserWithEmailPassword(c echo.Context) {
	var request authentication.UserLoginRequest
	helper.PanicIfError(c.Bind(&request))

	fmt.Printf("Request Email: %v\nRequest Password: %v\n", request.Email, request.Password)
}
