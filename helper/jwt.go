package helper

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	ApplicationName    = "Online Learning RESTful Api"
	ExpirationDuration = 48 * time.Hour
	JwtSigningMethod   = jwt.SigningMethodHS256
	JwtSignatureKey    = []byte("Secret Of Fikran Akbar")
)

func GenerateJWT(id uint, email string, username string) string {
	token := jwt.New(JwtSigningMethod)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["user_email"] = email
	claims["user_name"] = username
	claims["exp"] = ExpirationDuration
	t, err := token.SignedString(JwtSignatureKey)
	PanicIfError(err)
	return t
}
