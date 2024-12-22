package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type tokenService struct {
	secretKey     string
	tokenLifetime time.Duration
}

func NewTokenService(secretKey string, tokenLifetime time.Duration) TokenService {
	return &tokenService{
		secretKey:     secretKey,
		tokenLifetime: tokenLifetime,
	}
}

func (s *tokenService) GenerateAccessToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(s.tokenLifetime).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *tokenService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(s.secretKey), nil
	})
}
