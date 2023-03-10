package login_user

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"online-learning-restful-api/app/router"
	"online-learning-restful-api/di"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/authentication"
	"online-learning-restful-api/test"
	"regexp"
	"strings"
	"testing"
)

var (
	userLoginSuccessRequestBody                    = `{"email":"mollypotts@gmail.com","password":"Password"}`
	userLoginFailedEmptyEmailRequestBody           = `{"email":"","password":"Password"}`
	userLoginFailedWrongEmailOrPasswordRequestBody = `{"email":"mollypotts@gmail.com","password":"test"}`
)

func TestLoginUserSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	req := httptest.NewRequest(
		http.MethodPost,
		router.HostURLTest+router.UsersURLPath+router.LoginURLPath,
		strings.NewReader(userLoginSuccessRequestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse
	var userLoginResponse authentication.UserLoginResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &userLoginResponse)
	assert.Equal(t, "mollypotts@gmail.com", userLoginResponse.Email)
	assert.NotNil(t, userLoginResponse.Token)
}

func TestLoginUserFailedEmptyEmail(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	req := httptest.NewRequest(
		http.MethodPost,
		router.HostURLTest+router.UsersURLPath+router.LoginURLPath,
		strings.NewReader(userLoginFailedEmptyEmailRequestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusBadRequest, responseBody.Status)
	assert.Equal(t, test.MessageBadRequest, responseBody.Message)
	assert.Regexp(t, regexp.MustCompile(`(?i)validation`), responseBody.Data.(string))
}

func TestLoginUserFailedWrongEmailOrPassword(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	req := httptest.NewRequest(
		http.MethodPost,
		router.HostURLTest+router.UsersURLPath+router.LoginURLPath,
		strings.NewReader(userLoginFailedWrongEmailOrPasswordRequestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusUnauthorized, responseBody.Status)
	assert.Equal(t, test.MessageUnauthorized, responseBody.Message)
	assert.Regexp(t, regexp.MustCompile(`(?i)unauthorized`), responseBody.Data.(string))
}
