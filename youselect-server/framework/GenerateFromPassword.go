package framework

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(text *string) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*text), 3)
	if err != nil {
		return nil, err
	}

	encryptedPassword := string(hashedPassword)

	return &encryptedPassword, nil
}
