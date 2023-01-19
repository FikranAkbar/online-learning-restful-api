package get_detail_course_by_course_id

import (
	"encoding/json"
	"fmt"
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
	"regexp"
	"testing"
)

func TestGetDetailCourseSuccessWithAlreadyOwnedTrue(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseId := "/1"
	userIdQuery := "?userid=1"
	urlRoute := router.HostURL + router.CoursesAPIRoute + courseId + userIdQuery
	req := httptest.NewRequest(
		http.MethodGet,
		urlRoute,
		nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse
	var detailCourseResponse course.DetailCourseResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &detailCourseResponse)
	fmt.Println(detailCourseResponse)
	assert.Equal(t, "Intro to Entrepreneurship", detailCourseResponse.Name)
	assert.Equal(t, "Ray Robinson", detailCourseResponse.Expert.Name)
	assert.NotNil(t, detailCourseResponse.PhotoUrl)
	assert.NotNil(t, detailCourseResponse.Expert.Photo)
	assert.Equal(t, true, detailCourseResponse.AlreadyOwned)
}

func TestGetDetailCourseSuccessWithAlreadyOwnedFalse(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseId := "/1"
	userIdQuery := "?userid=1010101"
	urlRoute := router.HostURL + router.CoursesAPIRoute + courseId + userIdQuery
	req := httptest.NewRequest(
		http.MethodGet,
		urlRoute,
		nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse
	var detailCourseResponse course.DetailCourseResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &detailCourseResponse)
	fmt.Println(detailCourseResponse)
	assert.Equal(t, "Intro to Entrepreneurship", detailCourseResponse.Name)
	assert.Equal(t, "Ray Robinson", detailCourseResponse.Expert.Name)
	assert.NotNil(t, detailCourseResponse.PhotoUrl)
	assert.NotNil(t, detailCourseResponse.Expert.Photo)
	assert.Equal(t, false, detailCourseResponse.AlreadyOwned)
}

func TestGetDetailCourseFailedCourseNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseId := "/10101"
	userIdQuery := "?userid=1"
	urlRoute := router.HostURL + router.CoursesAPIRoute + courseId + userIdQuery
	req := httptest.NewRequest(
		http.MethodGet,
		urlRoute,
		nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusNotFound, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusNotFound, responseBody.Status)
	assert.Equal(t, test.MessageNotFound, responseBody.Message)
	assert.Regexp(t, regexp.MustCompile(`(?i)not found`), responseBody.Data.(string))
}
