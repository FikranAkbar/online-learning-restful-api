package get_industry_insights

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
	"testing"
)

func TestGetIndustryInsightsSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	urlRoute := router.HostURLTest +
		router.IndustryInsightsURLPath
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
	var industryInsightsResponse []industry_insight.IndustryInsightResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &industryInsightsResponse)
	assert.NotEqual(t, 0, len(industryInsightsResponse))
	assert.NotNil(t, industryInsightsResponse[0].TitleInsight)
	assert.NotNil(t, industryInsightsResponse[0].BodyContent)
	assert.NotNil(t, industryInsightsResponse[0].IdUserAuthor)
}
