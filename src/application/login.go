package application

import (
	"context"
	"time"

	"auth/src/domain"
	"auth/src/infrastructure/token"
)

type loginApplication struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginApplication(userRepository domain.UserRepository, timeout time.Duration) domain.LoginApplication {
	return &loginApplication{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
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