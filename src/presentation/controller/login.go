package controller

import (
	"auth/src/application"
	"auth/src/domain"

	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (LoginResponse, error) {
	if not (domain.IsValidEmail(email) and domain.IsValidPassword(password)) {
		return utils.ErrorResponse(utils.SCHEMA_NOT_MATCH, "params are not valid")
	}

	ur :=  repository.SetDefault()
	return applications.Login(ur, env.SystemUid, email, password)
}