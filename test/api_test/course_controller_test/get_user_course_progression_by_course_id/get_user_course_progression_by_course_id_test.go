package get_user_course_progression_by_course_id

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

var currentJWTToken = ""

func TestMain(m *testing.M) {
	currentJWTToken = test.LoginUserFirst(`{"email":"mollypotts@gmail.com","password":"Password"}`)
	code := m.Run()
	os.Exit(code)
}

func TestGetUserCourseProgressionByCourseIdSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseId := "/1"
	urlRoute := router.HostURLTest + router.CoursesAPIRoute + courseId + router.CourseProgressionsAPIRoute
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
	var userCourseProgressionResponse course.UserCourseProgressionResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &userCourseProgressionResponse)
	assert.Equal(t, uint(1), userCourseProgressionResponse.UserId)
	assert.Equal(t, uint(1), userCourseProgressionResponse.CourseId)
	assert.Equal(t, uint(1), userCourseProgressionResponse.LastUnlockedModule)
}

func TestGetUserCourseFailedNotOwnedTheAccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseId := "/101"
	urlRoute := router.HostURLTest + router.CoursesAPIRoute + courseId + router.CourseProgressionsAPIRoute
	req := httptest.NewRequest(
		http.MethodGet,
		urlRoute,
		nil)
	req.Header.Set("Authorization", "Bearer "+currentJWTToken)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusForbidden, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusForbidden, responseBody.Status)
	assert.Equal(t, test.MessageForbidden, responseBody.Message)
	assert.Regexp(t, regexp.MustCompile(`(?i)not owned the access`), responseBody.Data.(string))
}

func TestGetUserCourseFailedInvalidToken(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseId := "/101"
	urlRoute := router.HostURLTest + router.CoursesAPIRoute + courseId + router.CourseProgressionsAPIRoute
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
	assert.Regexp(t, regexp.MustCompile(`(?i)invalid token`), responseBody.Data.(string))
}
