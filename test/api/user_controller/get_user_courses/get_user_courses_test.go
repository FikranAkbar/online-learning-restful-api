package get_user_courses

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
	"online-learning-restful-api/model/web/course"
	"online-learning-restful-api/test"
	"os"
	"regexp"
	"testing"
)

var (
	queryCourseStatus      = "?type=ongoing"
	queryCourseStatusEmpty = "?type="
)

var currentJWTToken = ""

func TestMain(m *testing.M) {
	currentJWTToken = test.LoginUserFirst(`{"email":"mollypotts@gmail.com","password":"Password"}`)
	code := m.Run()
	os.Exit(code)
}

func TestGetUserCoursesSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	urlRoute := router.HostURLTest +
		router.UsersURLPath +
		router.CourseURLPath +
		queryCourseStatus

	req := httptest.NewRequest(
		http.MethodGet,
		urlRoute,
		nil,
	)
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse
	var coursesResponse []course.ShortCourseResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &coursesResponse)
	assert.NotNil(t, coursesResponse[0].Name)
	assert.NotNil(t, coursesResponse[0].ExpertName)
	assert.NotNil(t, coursesResponse[0].PhotoUrl)
}

func TestGetUserCoursesFailedEmptyQuery(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	urlRoute := router.HostURLTest +
		router.UsersURLPath +
		router.CourseURLPath +
		queryCourseStatusEmpty

	req := httptest.NewRequest(
		http.MethodGet,
		urlRoute,
		nil,
	)
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
	assert.Regexp(t, regexp.MustCompile(`(?i)query`), responseBody.Data)
	assert.Regexp(t, regexp.MustCompile(`(?i)empty`), responseBody.Data)
}

func TestGetUserCoursesFailedInvalidToken(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	urlRoute := router.HostURLTest +
		router.UsersURLPath +
		router.CourseURLPath +
		queryCourseStatusEmpty

	req := httptest.NewRequest(
		http.MethodGet,
		urlRoute,
		nil,
	)
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
