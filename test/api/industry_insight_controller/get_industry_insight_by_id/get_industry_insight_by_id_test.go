package get_industry_insight_by_id

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
	"online-learning-restful-api/model/web/industry_insight"
	"online-learning-restful-api/test"
	"regexp"
	"testing"
)

func TestGetIndustryInsightByIdSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	industryInsightIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.IndustryInsightsURLPath +
		industryInsightIdPath
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
	var industryInsightResponse industry_insight.IndustryInsightResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &industryInsightResponse)
	assert.NotNil(t, industryInsightResponse.TitleInsight)
	assert.NotNil(t, industryInsightResponse.BodyContent)
	assert.NotNil(t, industryInsightResponse.IdUserAuthor)
	assert.NotNil(t, industryInsightResponse.InsightAttachments)
}

func TestGetIndustryInsightByIdFailedNotFound(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	industryInsightIdPath := "/101"
	urlRoute := router.HostURLTest +
		router.IndustryInsightsURLPath +
		industryInsightIdPath
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
