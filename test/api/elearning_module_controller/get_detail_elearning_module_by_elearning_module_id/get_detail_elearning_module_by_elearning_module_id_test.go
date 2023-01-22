package get_detail_elearning_module_by_elearning_module_id

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"online-learning-restful-api/app/router"
	"online-learning-restful-api/di"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/elearning_module"
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

func TestGetDetailElearningModuleSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	elearningModuleIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.LearnURLPath +
		router.ElearningModuleURLPath +
		elearningModuleIdPath
	req := httptest.NewRequest(
		http.MethodGet,
		urlRoute,
		nil)
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse
	var detailElearningModulesResponse elearning_module.DetailElearningModuleResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &detailElearningModulesResponse)
	assert.NotNil(t, detailElearningModulesResponse)
	assert.NotNil(t, detailElearningModulesResponse.Title)
	assert.NotNil(t, detailElearningModulesResponse.Description)
	assert.NotNil(t, detailElearningModulesResponse.Quiz.Question)
	assert.NotNil(t, detailElearningModulesResponse.Video.Url)
	assert.NotNil(t, detailElearningModulesResponse.Video.Name)
}

func TestGetDetailElearningModuleFailedCourseNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/101"
	elearningModuleIdPath := "/101"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.LearnURLPath +
		router.ElearningModuleURLPath +
		elearningModuleIdPath
	req := httptest.NewRequest(
		http.MethodGet,
		urlRoute,
		nil)
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

func TestGetDetailElearningModuleFailedElearningModuleNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	elearningModuleIdPath := "/101"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.LearnURLPath +
		router.ElearningModuleURLPath +
		elearningModuleIdPath
	req := httptest.NewRequest(
		http.MethodGet,
		urlRoute,
		nil)
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
