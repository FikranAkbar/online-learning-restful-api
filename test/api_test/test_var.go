package api_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"io"
	"net/http"
	"net/http/httptest"
	"online-learning-restful-api/di"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/authentication"
	"strings"
)

// HostURL Host URL for test
var (
	HostURL = "http://localhost:8000/api"
)

// Users routes
var (
	UserLoginAPIRoute    = HostURL + "/users/login"
	UserLogoutAPIRoute   = HostURL + "/users/logout"
	UserRegisterAPIRoute = HostURL + "/users/register"
)

// messages
var (
	MessageOk           = "OK"
	MessageBadRequest   = "BAD REQUEST"
	MessageUnauthorized = "UNAUTHORIZED"
	MessageNotFound     = "NOT FOUND"
)

func LoginUserFirst() string {
	e := di.InitializedEchoServerForTest()

	req := httptest.NewRequest(http.MethodPost, UserLoginAPIRoute, strings.NewReader(`{email:"mollypotts@gmail.com", password:"Password"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse
	var userLoginResponse authentication.UserLoginResponse

	_ = json.Unmarshal(body, &responseBody)
	_ = mapstructure.Decode(responseBody.Data, &userLoginResponse)

	return userLoginResponse.Token
}
