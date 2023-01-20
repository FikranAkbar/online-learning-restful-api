package logout_user

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"online-learning-restful-api/app/router"
	"online-learning-restful-api/di"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/test"
	"os"
	"regexp"
	"testing"
)

var currentJWTToken = ""

func TestMain(m *testing.M) {
	currentJWTToken = test.LoginUserFirst(`{"email":"mollypotts@gmail.com","password":"Password"}`)
	code := m.Run()
	os.Exit(code)
}

func TestLogoutUserSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	req := httptest.NewRequest(
		http.MethodPost,
		router.HostURLTest+router.UsersURLPath+router.LogoutURLPath,
		nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+currentJWTToken)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.Regexp(t, regexp.MustCompile(`(?i)token valid`), responseBody.Data)
}

func TestLogoutUserFailed(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	req := httptest.NewRequest(
		http.MethodPost,
		router.HostURLTest+router.UsersURLPath+router.LogoutURLPath,
		nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusUnauthorized, responseBody.Status)
	assert.Equal(t, test.MessageUnauthorized, responseBody.Message)
	assert.Regexp(t, regexp.MustCompile(`(?i)invalid token`), responseBody.Data)
}
