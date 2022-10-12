package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"praktikum/constants"
	"time"
)

func CreateToken(userId uint, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userId
	claims["name"] = name
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
