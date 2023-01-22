package get_expert_detail_by_id

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
	"online-learning-restful-api/model/web/expert"
	"online-learning-restful-api/test"
	"regexp"
	"testing"
)

func TestGetExpertDetailByIdSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	expertIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.ExpertsURLPath +
		expertIdPath
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
	var expertDetailResponse expert.DetailExpertResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &expertDetailResponse)
	assert.NotNil(t, expertDetailResponse.ExpertId)
	assert.NotNil(t, expertDetailResponse.Name)
}

func TestGetExpertDetailByIdFailedNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	expertIdPath := "/101"
	urlRoute := router.HostURLTest +
		router.ExpertsURLPath +
		expertIdPath
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
	assert.Regexp(t, regexp.MustCompile(`(?i)not found`), responseBody.Data)
}
