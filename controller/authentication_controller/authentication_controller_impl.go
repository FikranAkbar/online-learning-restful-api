package authentication_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
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

func (controller *AuthenticationControllerImpl) LoginUserWithEmailPassword(c echo.Context) error {
	var userLoginRequest authentication.UserLoginRequest
	err := c.Bind(&userLoginRequest)
	helper.PanicIfError(err)

	userLoginResponse := controller.AuthenticationService.LoginUserByEmailPassword(c.Request().Context(), userLoginRequest)

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    userLoginResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *AuthenticationControllerImpl) RegisterUserWithEmailPassword(c echo.Context) error {
	var userRegisterRequest authentication.UserRegisterRequest
	err := c.Bind(&userRegisterRequest)
	helper.PanicIfError(err)

	userRegisterResponse := controller.AuthenticationService.RegisterUserByEmailPassword(c.Request().Context(), userRegisterRequest)

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    userRegisterResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *AuthenticationControllerImpl) LogoutUser(c echo.Context) error {
	ctx := helper.ExtractUserInfoFromJwtToContext(c)
	text := controller.AuthenticationService.LogoutUser(ctx)

	apiResponse := web.APIResponse{
		Status:  200,
		Message: "OK",
		Data:    text,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
