package common

import (
	"GinProject/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("secret-key")

type Claims struct {
	jwt.RegisteredClaims
	UserId uint
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "eric",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
