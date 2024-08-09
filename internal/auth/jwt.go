package auth

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, id int) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id, 
		"exp": expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", errors.New("error creating token")
	}
	return tokenString, nil
}