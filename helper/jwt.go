package helper

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtSecret = []byte("Secret Of Fikran Akbar")

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, _ := token.SignedString(jwtSecret)
	return t
}
