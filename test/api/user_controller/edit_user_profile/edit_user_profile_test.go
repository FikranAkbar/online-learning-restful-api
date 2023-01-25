package edit_user_profile

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"online-learning-restful-api/app/router"
	"online-learning-restful-api/di"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/user"
	"online-learning-restful-api/test"
	"os"
	"regexp"
	"strings"
	"testing"
)

var (
	editUserProfileRequest = user.EditProfileRequest{
		Name:         "Lala Potts",
		Phone:        "081231231",
		Gender:       "Female",
		Photo:        "https://www.google.com",
		ProvinceName: "DAERAH ISTIMEWA YOGYAKARTA",
		CityName:     "SLEMAN",
	}
)

var currentJWTToken = ""

func TestMain(m *testing.M) {
	currentJWTToken = test.LoginUserFirst(`{"email":"mollypotts@gmail.com","password":"Password"}`)
	code := m.Run()
	os.Exit(code)
}

func TestEditUserProfileSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	urlRoute := router.HostURLTest +
		router.UsersURLPath +
		router.EditProfileURLPath

	editUserProfileRequestJSON, err := json.Marshal(editUserProfileRequest)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPut,
		urlRoute,
		strings.NewReader(string(editUserProfileRequestJSON)),
	)
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse
	var userEditProfileResponse user.DetailUserResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	fmt.Println(responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &userEditProfileResponse)
	assert.NotNil(t, userEditProfileResponse.Name)
	assert.NotNil(t, userEditProfileResponse.ProvinceId)
}

func TestEditUserProfileFailedInvalidToken(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	urlRoute := router.HostURLTest +
		router.UsersURLPath +
		router.EditProfileURLPath

	editUserProfileRequestJSON, err := json.Marshal(editUserProfileRequest)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPut,
		urlRoute,
		strings.NewReader(string(editUserProfileRequestJSON)),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusUnauthorized, responseBody.Status)
	assert.Equal(t, test.MessageUnauthorized, responseBody.Message)
	assert.Regexp(t, regexp.MustCompile(`(?i)invalid token`), responseBody.Data)
}

func TestEditUserProfileFailedEmptyRequest(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	urlRoute := router.HostURLTest +
		router.UsersURLPath +
		router.EditProfileURLPath

	req := httptest.NewRequest(
		http.MethodPut,
		urlRoute,
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusBadRequest, responseBody.Status)
	assert.Equal(t, test.MessageBadRequest, responseBody.Message)
	assert.Regexp(t, regexp.MustCompile(`(?i)empty`), responseBody.Data)
}
