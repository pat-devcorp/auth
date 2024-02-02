package controller

import (
	"auth/src/application"
	"auth/src/domain"

	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewLogin(email string, password string) (LoginResponse, error) {
	if not (domain.IsValidEmail(email) and domain.IsValidPassword(password)) {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid parameters"})
		return
	}

	user, err := application.LoginApplication.GetUserByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := application.LoginApplication.CreateAccessToken(&user, application.Env.AccessTokenSecret, application.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := application.LoginApplication.CreateRefreshToken(&user, application.Env.RefreshTokenSecret, application.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	return LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}