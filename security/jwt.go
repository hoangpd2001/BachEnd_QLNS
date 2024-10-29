package security

import (
	"BackEnd/mod/model"
	resUser "BackEnd/mod/model/model_user/res_user"
	"log"

	"time"

	"github.com/golang-jwt/jwt"
)

const SECRET_KEY = "hoang"

func GenToken(user resUser.ResSingin, roles []int) (string, error) {

	claims := &model.JwtCustomClaims{
		UserId: user.ID,
		Role:   roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 100).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return result, nil
}
func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := SECRET_KEY
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
