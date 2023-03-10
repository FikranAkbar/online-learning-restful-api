package user_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/user"
	"online-learning-restful-api/service/user_service"
	"strconv"
)

type UserControllerImpl struct {
	user_service.UserService
}

func NewUserControllerImpl(userService user_service.UserService) *UserControllerImpl {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) GetUserCourses(c echo.Context) error {
	courseStatus := c.QueryParam("type")

	coursesResponse := controller.UserService.GetUserCourses(c.Request().Context(), courseStatus)

	apiResponse := web.APIResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    coursesResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *UserControllerImpl) GetProvinces(c echo.Context) error {
	provincesResponse := controller.UserService.GetProvinces(c.Request().Context())

	apiResponse := web.APIResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    provincesResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *UserControllerImpl) GetCitiesByProvinceId(c echo.Context) error {
	provinceId, err := strconv.Atoi(c.Param("provinceId"))
	helper.PanicIfError(err)

	citiesResponse := controller.UserService.GetCitiesByProvinceId(c.Request().Context(), uint(provinceId))

	apiResponse := web.APIResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    citiesResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (controller *UserControllerImpl) EditUserProfile(c echo.Context) error {
	if c.Request().Body == http.NoBody {
		logError := "Cannot edit profile with empty request body"
		panic(exception.GenerateHTTPError(exception.BadRequest, logError))
	}

	var request user.EditProfileRequest
	err := c.Bind(&request)
	helper.PanicIfError(err)

	userResponse := controller.UserService.EditUserProfile(c.Request().Context(), request)

	apiResponse := web.APIResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    userResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
