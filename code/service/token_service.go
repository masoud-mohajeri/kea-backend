package service

import (
	"time"

	"github.com/masoud-mohajeri/kea-backend/dto"

	"github.com/masoud-mohajeri/kea-backend/config"
	"github.com/masoud-mohajeri/kea-backend/constants"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	CreateToken(mobile string, role constants.UserRole) (*dto.TokenDto, error)
	ParsToken(tokenString string) (*jwt.Token, error)
}

type tokenService struct{}

func NewTokenService() TokenService {
	return &tokenService{}
}

func (ts *tokenService) CreateToken(mobile string, role constants.UserRole) (*dto.TokenDto, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"mobile": mobile,
			"role":   role,
			"exp":    time.Now().Add(config.CommonConfig.AccessTokenDuration).Unix(),
		})

	tokenString, err := token.SignedString(config.CommonConfig.TokenSecret)
	if err != nil {
		return nil, err
	}

	return &dto.TokenDto{AccessToken: tokenString}, nil
}

func (ts *tokenService) ParsToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.CommonConfig.TokenSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
