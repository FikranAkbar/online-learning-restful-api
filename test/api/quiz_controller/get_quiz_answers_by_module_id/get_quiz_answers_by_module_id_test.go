package get_quiz_answers_by_module_id

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
	"online-learning-restful-api/model/web/quiz"
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

func TestGetQuizAnswersSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	elearningModuleIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.LearnURLPath +
		router.ElearningModuleURLPath +
		elearningModuleIdPath +
		router.QuizAnswersURLPath

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
	var quizAnswersResponse []quiz.DetailQuizAnswerResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &quizAnswersResponse)
	assert.NotNil(t, quizAnswersResponse)
	assert.NotNil(t, quizAnswersResponse[0].QuizAnswer)
	assert.NotNil(t, quizAnswersResponse[0].UserName)
}

func TestGetQuizAnswersFailedInvalidToken(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	elearningModuleIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.LearnURLPath +
		router.ElearningModuleURLPath +
		elearningModuleIdPath +
		router.QuizAnswersURLPath

	req := httptest.NewRequest(
		http.MethodGet,
		urlRoute,
		nil)
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

func TestGetQuizAnswersFailedCourseNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/101"
	elearningModuleIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.LearnURLPath +
		router.ElearningModuleURLPath +
		elearningModuleIdPath +
		router.QuizAnswersURLPath

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
	assert.Regexp(t, regexp.MustCompile(`(?i)course`), responseBody.Data)
	assert.Regexp(t, regexp.MustCompile(`(?i)not found`), responseBody.Data)
}

func TestGetQuizAnswersFailedModuleNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	elearningModuleIdPath := "/101"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.LearnURLPath +
		router.ElearningModuleURLPath +
		elearningModuleIdPath +
		router.QuizAnswersURLPath

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
	assert.Regexp(t, regexp.MustCompile(`(?i)module`), responseBody.Data)
	assert.Regexp(t, regexp.MustCompile(`(?i)not found`), responseBody.Data)
}
