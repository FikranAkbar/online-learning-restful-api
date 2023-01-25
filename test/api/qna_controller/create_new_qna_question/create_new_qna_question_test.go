package create_new_qna_question

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
	"online-learning-restful-api/model/web/qna"
	"online-learning-restful-api/test"
	"os"
	"regexp"
	"strings"
	"testing"
)

var qnaQuestionRequest = qna.CreateQnaQuestionRequest{Question: "Ini Pertanyaan saya 3 (?)"}

var currentJWTToken = ""

func TestMain(m *testing.M) {
	currentJWTToken = test.LoginUserFirst(`{"email":"mollypotts@gmail.com","password":"Password"}`)
	code := m.Run()
	os.Exit(code)
}

func TestCreateQnaQuestionSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.QnaQuestionsURLPath

	qnaQuestionRequestJSON, err := json.Marshal(qnaQuestionRequest)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		strings.NewReader(string(qnaQuestionRequestJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse
	var qnaQuestionResponse qna.DetailQnaQuestionResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &qnaQuestionResponse)
	assert.NotNil(t, qnaQuestionResponse)
	assert.NotNil(t, qnaQuestionResponse.UserName)
	assert.NotNil(t, qnaQuestionResponse.Question)
}

func TestCreateQnaQuestionFailedInvalidToken(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.QnaQuestionsURLPath

	qnaQuestionRequestJSON, err := json.Marshal(qnaQuestionRequest)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		strings.NewReader(string(qnaQuestionRequestJSON)))
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

func TestCreateQnaQuestionFailedCourseNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/101"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.QnaQuestionsURLPath

	qnaQuestionRequestJSON, err := json.Marshal(qnaQuestionRequest)
	helper.PanicIfError(err)

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		strings.NewReader(string(qnaQuestionRequestJSON)))
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
	assert.Regexp(t, regexp.MustCompile(`(?i)course`), responseBody.Data)
	assert.Regexp(t, regexp.MustCompile(`(?i)not found`), responseBody.Data)
}

func TestCreateQnaQuestionFailedEmptyRequest(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.QnaQuestionsURLPath

	req := httptest.NewRequest(
		http.MethodPost,
		urlRoute,
		nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
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
