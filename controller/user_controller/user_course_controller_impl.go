package user_controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/service/user_service"
)

type UserCourseControllerImpl struct {
	user_service.UserService
}

func NewUserCourseControllerImpl(userService user_service.UserService) *UserCourseControllerImpl {
	return &UserCourseControllerImpl{UserService: userService}
}

func (controller *UserCourseControllerImpl) GetUserCourses(c echo.Context) error {
	courseStatus := c.QueryParam("type")

	coursesResponse := controller.UserService.GetUserCourses(c.Request().Context(), courseStatus)

	apiResponse := web.APIResponse{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    coursesResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
