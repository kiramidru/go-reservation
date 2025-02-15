package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	User_type  string `json:"user_type"`
	ID         string `json:"id"`
	jwt.RegisteredClaims
}

type JwtRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
