package entities

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Username string `json:"name"`
	jwt.RegisteredClaims
}
