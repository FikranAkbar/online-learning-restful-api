package test

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"io"
	"net/http"
	"net/http/httptest"
	"online-learning-restful-api/app/router"
	"online-learning-restful-api/di"
	"online-learning-restful-api/model/web"
	"online-learning-restful-api/model/web/authentication"
	"strings"
)

// messages
var (
	MessageOk           = "OK"
	MessageBadRequest   = "BAD REQUEST"
	MessageUnauthorized = "UNAUTHORIZED"
	MessageForbidden    = "FORBIDDEN"
	MessageNotFound     = "NOT FOUND"
)

func LoginUserFirst(userCredJSON string) string {
	e := di.InitializedEchoServerForTest()

	req := httptest.NewRequest(
		http.MethodPost,
		router.HostURL+router.UsersURLPath+router.LoginURLPath,
		strings.NewReader(userCredJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	response := rec.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody web.APIResponse
	var userLoginResponse authentication.UserLoginResponse

	_ = json.Unmarshal(body, &responseBody)
	_ = mapstructure.Decode(responseBody.Data, &userLoginResponse)

	return userLoginResponse.Token
}
