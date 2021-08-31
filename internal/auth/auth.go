package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("it`s_hard_salt_1")

func CreateNewToken(value string) (string, error) {
	claims := &Claims{
		Login: value,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(bearerHeader string) (interface{}, error) {
	bearerSplit := strings.Split(bearerHeader, " ")
	if len(bearerSplit) < 2 {
		return nil, errors.New("Invalid header\n")
	}

	bearerToken := bearerSplit[1]
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized, err
		}
		return http.StatusUnauthorized, err
	}
	if !tkn.Valid {
		return http.StatusUnauthorized, errors.New("Not valid token\n")
	}

	return claims.Login, nil
}
