package db

import (
	"twitter/models"

	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.User, bool) {
	user, create, _ := CheckIsExistUser(email)
	if create == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
