package controller

import (
	"auth/src/application"
	"auth/src/domain"

	"golang.org/x/crypto/bcrypt"
)

func NewLogin(email string, password string) (LoginResponse, error) {
	if not (domain.IsValidEmail(email) and domain.IsValidPassword(password)) {
		return utils.ErrorResponse(utils.SCHEMA_NOT_MATCH, "error in format")
	}

	ur :=  repository.SetDefault()
	return applications.NewLogin(ur, email, password)
}