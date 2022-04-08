package utility

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func CompareHash(plainText, hashText string) error {
	plainTextInBytes := []byte(plainText)
	hashTextInBytes := []byte(hashText)
	err := bcrypt.CompareHashAndPassword(hashTextInBytes, plainTextInBytes)
	return err
}
func CheckPassword(password, hash string) bool {
	fmt.Println(password, hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}

func GenPasswordHash(password string) (string, error) {
	password = strings.TrimSpace(password)
	pwdByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pwdByte), nil
}
