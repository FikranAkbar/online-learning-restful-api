package get_overview_webinar_session_by_course_id

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
	"online-learning-restful-api/model/web/webinar_session"
	"online-learning-restful-api/test"
	"regexp"
	"testing"
)

func TestGetOverviewWebinarSessionSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.CoursesAPIRoute +
		courseIdPath +
		router.OverviewAPIRoute +
		router.WebinarSessionsAPIRoute
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
	var overviewWebinarSessionsResponse []webinar_session.OverviewWebinarSessionResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &overviewWebinarSessionsResponse)
	assert.NotNil(t, overviewWebinarSessionsResponse)
	assert.NotEqual(t, 0, len(overviewWebinarSessionsResponse))
	assert.NotNil(t, overviewWebinarSessionsResponse[0].CourseId)
	assert.NotNil(t, overviewWebinarSessionsResponse[0].WebinarSessionId)
}

func TestGetOverviewWebinarSessionFailedCourseNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	courseIdPath := "/101"
	urlRoute := router.HostURLTest +
		router.CoursesAPIRoute +
		courseIdPath +
		router.OverviewAPIRoute +
		router.WebinarSessionsAPIRoute
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
