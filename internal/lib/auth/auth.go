package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gitlab.ubrato.ru/ubrato/core/internal/config"
)

var (
	// errEmptySecret is returned when empty secret received in the settings.
	errEmptySecret = errors.New("authorizer empty JWT secret not allowed")
)

type Claims struct {
	jwt.RegisteredClaims
	Payload
}

type TokenAuthorizer struct {
	cfg config.JWTSettings
}

// NewTokenAuthorizer returns TokenAuthorizer.
func NewTokenAuthorizer(settings config.JWTSettings) (*TokenAuthorizer, error) {
	if settings.Secret == "" {
		return nil, errEmptySecret
	}

	return &TokenAuthorizer{cfg: settings}, nil
}

func (ta *TokenAuthorizer) GetRefreshTokenDurationLifetime() time.Duration {
	return ta.cfg.Lifetime.Refresh
}

func (ta *TokenAuthorizer) GenerateToken(payload Payload) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ta.cfg.Lifetime.Access)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Payload: payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(ta.cfg.Secret))
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return ss, nil
}
