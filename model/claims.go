package model

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	UserId int
	Role string
	jwt.StandardClaims
}