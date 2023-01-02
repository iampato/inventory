package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword string, password string) (bool, error) {
	// convert string to []byte
	passwordBytes := []byte(hashedPassword)
	// Check if the hashed password is correct
	err := bcrypt.CompareHashAndPassword(passwordBytes, []byte(password))
	if err != nil {
		fmt.Printf("Error checking password: %v", err)
		return false, err
	}
	return true, nil
}
