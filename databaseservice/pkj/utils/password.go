package utils

import (
	

	"golang.org/x/crypto/bcrypt"
)

func validatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	return true
}

func hashPassword(password string) string {
	hashedpassword,err:=bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		println(err)
	}
	return string(hashedpassword)
}
func comparePassword(password string, hashedpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}