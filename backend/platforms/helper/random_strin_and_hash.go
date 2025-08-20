package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CompareHash this function compares the hash with the string and returns a boolean value accordingly.
func CompareHash(hash, password string) bool {
	if eror := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); eror != nil {
		return false
	}
	return true
}
