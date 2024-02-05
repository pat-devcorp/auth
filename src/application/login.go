package application

import (
	"context"
	"time"

	"auth/src/domain"
	"auth/src/utils"
	"auth/src/infrastructure/token"
)

type loginApplication struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (lu *loginApplication) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginApplication) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return token.CreateAccessToken(user, secret, expiry)
}

func (lu *loginApplication) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return token.CreateRefreshToken(user, secret, expiry)
}

func (lu *loginApplication) NewLogin(c context.Context, email string, password string) {
	user, err := GetUserByEmail(c, email)
	if err != nil {
		return utils.ErrorResponse(utils.UNAUTHORIZED, "User not found with the given email")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return utils.ErrorResponse(utils.WRONG_PASSWORD, "Invalid credentials")
	}

	accessToken, err := CreateAccessToken(&user, application.Env.AccessTokenSecret, application.Env.AccessTokenExpiryHour)
	if err != nil {
		return utils.ErrorResponse(utils.LOGIC_CRASH, err.Error())
	}

	refreshToken, err := CreateRefreshToken(&user, application.Env.RefreshTokenSecret, application.Env.RefreshTokenExpiryHour)
	if err != nil {
		return utils.ErrorResponse(utils.LOGIC_CRASH, err.Error())
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return LoginResponse
}