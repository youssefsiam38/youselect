package framework

import (
	"golang.org/x/crypto/bcrypt"
)

// CompareHashandPassword compares the hash and the password but in strings
func CompareHashandPassword(hashedText, text string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(text))

	if err != nil {
		return err
	}

	return nil
}
