package authentication_service

import (
	"context"
	"online-learning-restful-api/model/web/authentication"
)

type AuthenticationService interface {
	LoginUserByEmailPassword(ctx context.Context, request authentication.UserLoginRequest)
}
