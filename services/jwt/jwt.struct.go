package jwt

import "github.com/golang-jwt/jwt/v5"

type CustomJwtClaims struct {
	ID uint `json:"user_id"`
	jwt.RegisteredClaims
}
