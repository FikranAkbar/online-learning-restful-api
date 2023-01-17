package register_user_test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"online-learning-restful-api/app/database"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/di"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/authentication"
	"online-learning-restful-api/test/api_test"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var (
	userRegisterSuccessRequestBody = `{
  		"email": "fikranakbar756@gmail.com",
  		"password": "password",
  		"name": "Fikran Akbar",
  		"birth_date": "2000-07-01",
  		"gender": "Male"
	}`
	userRegisterFailedFieldNotComplete = `{
  		"email": "fikranakbar756@gmail.com",
  		"password": "password"
	}`
	userRegisterFailedEmailAlreadyExist = `{
  		"email": "mollypotts@gmail.com",
  		"password": "password",
  		"name": "Fikran Akbar",
  		"birth_date": "2000-07-01",
  		"gender": "Male"
	}`
)

func TestMain(m *testing.M) {
	code := m.Run()

	db := database.NewTestDB()

	var accountEntity entity.MasterAccount
	var userEntity entity.MasterUser

	tx := db.Begin()
	defer helper.CommitOrRollback(tx)

	tx.Where("email = ?", "fikranakbar756@gmail.com").First(&accountEntity)
	tx.Where("id = ?", strconv.Itoa(int(accountEntity.ID))).Delete(&userEntity)
	tx.Where("id = ?", strconv.Itoa(int(accountEntity.ID))).Delete(&accountEntity)

	os.Exit(code)
}

func TestTableRegisterUser(t *testing.T) {
	e := di.InitializedEchoServerForTest()
	t.Run("Register User Success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, api_test.UserRegisterAPIRoute, strings.NewReader(userRegisterSuccessRequestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		response := rec.Result()
		assert.Equal(t, http.StatusOK, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody web.APIResponse
		var userRegisterResponse authentication.UserRegisterResponse

		_ = json.Unmarshal(body, &responseBody)
		assert.Equal(t, http.StatusOK, responseBody.Status)
		assert.Equal(t, api_test.MessageOk, responseBody.Message)
		assert.NotNil(t, responseBody.Data)

		_ = mapstructure.Decode(responseBody.Data, &userRegisterResponse)
		assert.Equal(t, "fikranakbar756@gmail.com", userRegisterResponse.Email)
		assert.Equal(t, "Fikran Akbar", userRegisterResponse.Name)
	})
	t.Run("Register User Failed Field Not Complete", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, api_test.UserRegisterAPIRoute, strings.NewReader(userRegisterFailedFieldNotComplete))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		response := rec.Result()
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody web.APIResponse

		_ = json.Unmarshal(body, &responseBody)
		assert.Equal(t, http.StatusBadRequest, responseBody.Status)
		assert.Equal(t, api_test.MessageBadRequest, responseBody.Message)
		assert.Regexp(t, regexp.MustCompile(`(?i)bad request`), responseBody.Data)
	})
	t.Run("Register User Failed Email Already Exist", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, api_test.UserRegisterAPIRoute, strings.NewReader(userRegisterFailedEmailAlreadyExist))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		response := rec.Result()
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)

		body, _ := io.ReadAll(response.Body)
		var responseBody web.APIResponse

		_ = json.Unmarshal(body, &responseBody)
		assert.Equal(t, http.StatusBadRequest, responseBody.Status)
		assert.Equal(t, api_test.MessageBadRequest, responseBody.Message)
		assert.Regexp(t, regexp.MustCompile(`(?i)already exist`), responseBody.Data)
	})
}