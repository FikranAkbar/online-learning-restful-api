package create_new_quiz_answer

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
	"online-learning-restful-api/model/web/quiz"
	"online-learning-restful-api/test"
	"os"
	"regexp"
	"strings"
	"testing"
)

var quizAnswerRequestSuccess = quiz.ShortQuizAnswerRequest{
	Answer: "Test",
}

var quizAnswerRequestFailed = quiz.ShortQuizAnswerRequest{
	Answer: "",
}

var currentJWTToken = ""

func TestMain(m *testing.M) {
	currentJWTToken = test.LoginUserFirst(`{"email":"mollypotts@gmail.com","password":"Password"}`)
	code := m.Run()
	os.Exit(code)
}

func TestCreateNewQuizAnswerSuccess(t *testing.T) {
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

	requestJSON, err := json.Marshal(quizAnswerRequestSuccess)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		strings.NewReader(string(requestJSON)))
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse
	var quizAnswersResponse quiz.ShortQuizAnswerResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &quizAnswersResponse)
	assert.NotNil(t, quizAnswersResponse)
	assert.NotNil(t, quizAnswersResponse.QuizAnswer)
	assert.NotNil(t, quizAnswersResponse.UserId)
	assert.NotNil(t, quizAnswersResponse.QuizId)
}

func TestCreateNewQuizAnswerFailedInvalidToken(t *testing.T) {
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

	requestJSON, err := json.Marshal(quizAnswerRequestFailed)
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

func TestCreateNewQuizAnswerFailedEmptyRequest(t *testing.T) {
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

	requestJSON, err := json.Marshal(quizAnswerRequestFailed)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		strings.NewReader(string(requestJSON)))
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
	assert.Regexp(t, regexp.MustCompile(`(?i)required`), responseBody.Data)
}

func TestCreateNewQuizAnswerCourseNotFound(t *testing.T) {
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

	requestJSON, err := json.Marshal(quizAnswerRequestSuccess)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		strings.NewReader(string(requestJSON)))
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
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
	assert.Regexp(t, regexp.MustCompile("(?i)not found"), responseBody.Data)
}

func TestCreateNewQuizAnswerModuleNotFound(t *testing.T) {
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

	requestJSON, err := json.Marshal(quizAnswerRequestSuccess)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		strings.NewReader(string(requestJSON)))
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
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
	assert.Regexp(t, regexp.MustCompile("(?i)not found"), responseBody.Data)
}
