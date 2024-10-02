package security

import (
	"BackEnd/mod/model"
	resUser "BackEnd/mod/model/model_user/res_user"
	"time"

	"github.com/golang-jwt/jwt"
)

const SECRET_KEY = "hoang"

func GenToken(user resUser.ResUser) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return result, nil
}
func GenTokenUserFull(user resUser.ResUserFull) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return result, nil
}
