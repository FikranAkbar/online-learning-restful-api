package get_cities_by_province_id

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
	"online-learning-restful-api/model/web/city"
	"online-learning-restful-api/test"
	"os"
	"testing"
)

var currentJWTToken = ""

func TestMain(m *testing.M) {
	currentJWTToken = test.LoginUserFirst(`{"email":"mollypotts@gmail.com","password":"Password"}`)
	code := m.Run()
	os.Exit(code)
}

func TestGetCitiesByProvinceIdSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	provinceIdPath := "/1"
	urlRoute := router.HostURLTest +
		router.UsersURLPath +
		router.EditProfileURLPath +
		router.ProvincesURLPath +
		provinceIdPath +
		router.CitiesURLPath

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
	var citiesResponse []city.ShortCityResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &citiesResponse)
	assert.NotNil(t, citiesResponse[0].Id)
	assert.NotNil(t, citiesResponse[0].CityName)
	assert.NotNil(t, citiesResponse[0].ProvinceId)
}
