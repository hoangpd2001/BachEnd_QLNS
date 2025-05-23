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
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
	fmt.Println(hash)
	fmt.Println(password)
		fmt.Println(err)
		return false

	}
	fmt.Println(hash)
	fmt.Println(password)
	return true
}
