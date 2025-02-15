package usecase

import (
	"carbon/go-reservation/domain"
	"carbon/go-reservation/internal/tokenutil"
	"context"
	"time"
)

type signupUsecase struct {
    userRepository domain.UserRepository
    contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
    return &signupUsecase {
        userRepository: userRepository,
        contextTimeout: timeout,
    }
}

func (su *signupUsecase) Create(c context.Contextm, user *domain.User) error {
    ctx, cancel := context.WithTimeout(c, su.contextTimeout)
    defer cancel()
    return su.userRepository.
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
    ctx, cancel := context.WithTimeout(c, su.contextTimeout)
    defer cancel()
    return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
    return tokenutil.CreateAccessToken(user, secret, expiry)
}
func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
    return tokenutil.CreateRefreshToken(user, secret, expiry)
}
