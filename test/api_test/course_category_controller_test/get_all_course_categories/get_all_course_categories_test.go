package get_all_course_categories

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"online-learning-restful-api/di"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/test"
	"testing"
)

func TestGetAllCourseCategoriesSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	req := httptest.NewRequest(http.MethodGet, test.GetAllCourseCategoriesAPIRoute, nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()
	assert.Equal(t, http.StatusOK, rec.Code)

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)
}
