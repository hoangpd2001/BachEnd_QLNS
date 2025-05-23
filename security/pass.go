package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashingPasswordFunc(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
func CheckPasswordHashFunc(password string, hash string) bool {
	fmt.Println(HashingPasswordFunc("123456789a"))
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false

	}
	fmt.Println(hash)
	fmt.Println(password)
	return true
}
