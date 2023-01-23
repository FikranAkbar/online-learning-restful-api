package get_provinces

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
	"online-learning-restful-api/model/web/province"
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

func TestGetProvincesSuccess(t *testing.T) {
	e := di.InitializedEchoServerForTest()

	urlRoute := router.HostURLTest +
		router.UsersURLPath +
		router.EditProfileURLPath +
		router.ProvincesURLPath

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
	var provincesResponse []province.ShortProvinceResponse

	_ = json.Unmarshal(body, &responseBody)
	assert.Equal(t, http.StatusOK, responseBody.Status)
	assert.Equal(t, test.MessageOk, responseBody.Message)
	assert.NotNil(t, responseBody.Data)

	_ = mapstructure.Decode(responseBody.Data, &provincesResponse)
	assert.NotNil(t, provincesResponse[0].Id)
	assert.NotNil(t, provincesResponse[0].ProvinceName)
}
