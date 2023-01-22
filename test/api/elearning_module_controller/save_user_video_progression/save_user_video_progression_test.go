package save_user_video_progression

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
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/video"
	"online-learning-restful-api/test"
	"os"
	"regexp"
	"strings"
	"testing"
)

var userVideoProgressionRequest = video.UserVideoProgressionRequest{
	Progression: 10000,
	IsComplete:  true,
}

var currentJWTToken = ""

func TestMain(m *testing.M) {
	currentJWTToken = test.LoginUserFirst(`{"email":"mollypotts@gmail.com","password":"Password"}`)
	code := m.Run()
	os.Exit(code)
}

func TestSaveUserVideoProgressionSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	elearningModuleIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.LearnURLPath +
		router.ElearningModuleURLPath +
		elearningModuleIdPath +
		router.SaveVideoProgressionURLPath

	requestJSON, err := json.Marshal(userVideoProgressionRequest)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		strings.NewReader(string(requestJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse
	var userVideoProgressionResponse video.UserVideoProgressionResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &userVideoProgressionResponse)
	assert.Equal(t, userVideoProgressionRequest.Progression, userVideoProgressionResponse.Progressions)
	assert.Equal(t, userVideoProgressionRequest.IsComplete, userVideoProgressionResponse.IsComplete)
}

func TestSaveUserVideoProgressionFailedInvalidToken(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	elearningModuleIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.LearnURLPath +
		router.ElearningModuleURLPath +
		elearningModuleIdPath +
		router.SaveVideoProgressionURLPath

	requestJSON, err := json.Marshal(userVideoProgressionRequest)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		strings.NewReader(string(requestJSON)))
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

func TestSaveUserVideoProgressionFailedCourseNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/101"
	elearningModuleIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.LearnURLPath +
		router.ElearningModuleURLPath +
		elearningModuleIdPath +
		router.SaveVideoProgressionURLPath

	requestJSON, err := json.Marshal(userVideoProgressionRequest)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		strings.NewReader(string(requestJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusNotFound, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusNotFound, responseBody.Status)
	assert.Equal(t, test.MessageNotFound, responseBody.Message)
	assert.Regexp(t, regexp.MustCompile(`(?i)course`), responseBody.Data.(string))
	assert.Regexp(t, regexp.MustCompile(`(?i)not found`), responseBody.Data.(string))
}

func TestSaveUserVideoProgressionFailedModuleNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	elearningModuleIdPath := "/101"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.LearnURLPath +
		router.ElearningModuleURLPath +
		elearningModuleIdPath +
		router.SaveVideoProgressionURLPath

	requestJSON, err := json.Marshal(userVideoProgressionRequest)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		strings.NewReader(string(requestJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusNotFound, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusNotFound, responseBody.Status)
	assert.Equal(t, test.MessageNotFound, responseBody.Message)
	assert.Regexp(t, regexp.MustCompile(`(?i)module`), responseBody.Data.(string))
	assert.Regexp(t, regexp.MustCompile(`(?i)not found`), responseBody.Data.(string))
}
