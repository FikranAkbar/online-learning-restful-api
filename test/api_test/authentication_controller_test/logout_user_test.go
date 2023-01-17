package authentication_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"online-learning-restful-api/di"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/test/api_test"
	"os"
	"regexp"
	"testing"
)

var currentJWTToken = ""

func TestMain(m *testing.M) {
	currentJWTToken = api_test.LoginUserFirst()
	code := m.Run()
	os.Exit(code)
}

func TestLogoutUserSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	req := httptest.NewRequest(http.MethodPost, api_test.UserLogoutAPIRoute, nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+currentJWTToken)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, api_test.MessageOk, responseBody.Message)
	assert.Regexp(t, regexp.MustCompile(`(?i)token valid`), responseBody.Data)
}