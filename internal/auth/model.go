package auth

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Login string `json:"id"`
	jwt.StandardClaims
}
