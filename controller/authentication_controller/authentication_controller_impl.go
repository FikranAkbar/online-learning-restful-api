package authentication_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/model/web"
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

func (controller *AuthenticationControllerImpl) LoginUserWithEmailPassword(c echo.Context) error {
	//var request authentication.UserLoginRequest
	//helper.PanicIfError(c.Bind(&request))
	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data: struct {
			Email string
			Name  string
			Token string
		}{
			Email: "test@gmail.com",
			Name:  "test",
			Token: "yJhbGc",
		},
	}

	//fmt.Printf("Request Email: %v\nRequest Password: %v\n", request.Email, request.Password)

	return c.JSON(http.StatusOK, apiResponse)
}
