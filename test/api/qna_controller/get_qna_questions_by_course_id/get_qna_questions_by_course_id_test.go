package get_qna_questions_by_course_id

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
	"online-learning-restful-api/model/web/qna"
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

func TestGetQnaQuestionsSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.QnaQuestionsURLPath

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
	var qnaQuestionsResponse []qna.DetailQnaQuestionResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &qnaQuestionsResponse)
	assert.NotNil(t, qnaQuestionsResponse)
	assert.NotNil(t, qnaQuestionsResponse[0].UserName)
	assert.NotNil(t, qnaQuestionsResponse[0].Question)
}

func TestGetQnaQuestionsFailedInvalidToken(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.QnaQuestionsURLPath

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

func TestGetQnaQuestionsFailedCourseNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/101"
	urlRoute := router.HostURLTest +
		router.CourseURLPath +
		courseIdPath +
		router.QnaQuestionsURLPath

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
