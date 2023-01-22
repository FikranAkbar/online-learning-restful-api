package get_expert_courses_by_id

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
	"testing"
)

func TestGetExpertCoursesByIdSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	expertIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.ExpertsURLPath +
		expertIdPath +
		router.CourseURLPath
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
	var coursesResponse []course.ShortCourseResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &coursesResponse)
	assert.NotNil(t, coursesResponse)
	assert.NotNil(t, coursesResponse[0].Name)
	assert.NotNil(t, coursesResponse[0].ExpertName)
}
