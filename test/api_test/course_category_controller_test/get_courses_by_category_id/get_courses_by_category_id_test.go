package get_courses_by_category_id

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

func TestGetCoursesByCategoryIdSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	categoryId := "/1"
	urlRoute := router.HostURLTest + router.CoursesAPIRoute + router.CategoriesAPIRoute + categoryId
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
	assert.Equal(t, "Intro to Entrepreneurship", coursesResponse[0].Name)
	assert.Equal(t, "Ray Robinson", coursesResponse[0].ExpertName)
	assert.Equal(t, "https://images.unsplash.com/photo-1432888498266-38ffec3eaf0a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1174&q=80",
		coursesResponse[0].PhotoUrl)
}
