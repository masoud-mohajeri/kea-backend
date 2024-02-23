package service

import (
	"errors"
	"time"

	"github.com/masoud-mohajeri/kea-backend/dto"
	"github.com/masoud-mohajeri/kea-backend/entity"

	"github.com/masoud-mohajeri/kea-backend/config"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	CreateToken(user *entity.User) (*dto.TokenDto, error)
	ExtractToken(tokenString string) (string, error)
}

type tokenService struct {
}

func NewTokenService() TokenService {
	return &tokenService{}
}

func (ts *tokenService) CreateToken(user *entity.User) (*dto.TokenDto, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"mobile": user.Mobile,
			"role":   user.Role,
			"exp":    time.Now().Add(config.CommonConfig.AccessTokenDuration).Unix(),
		})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"mobile": user.Mobile,
			"exp":    time.Now().Add(config.CommonConfig.RefreshTokenDuration).Unix(),
		})

	accessTokenString, errAcc := accessToken.SignedString(config.CommonConfig.TokenSecret)
	refreshTokenString, errRef := refreshToken.SignedString(config.CommonConfig.TokenSecret)
	if errAcc != nil {
		return nil, errAcc
	}
	if errRef != nil {
		return nil, errRef
	}

	return &dto.TokenDto{AccessToken: accessTokenString, RefreshToken: refreshTokenString}, nil
}

func (ts *tokenService) parsToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.CommonConfig.TokenSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (ts *tokenService) ExtractToken(tokenString string) (string, error) {
	token, errT := ts.parsToken(tokenString)
	if errT != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	exp, _ := token.Claims.GetExpirationTime()
	if exp.Before(time.Now()) {
		return "", errors.New("token expired")
	}

	claims := token.Claims.(jwt.MapClaims)
	mobile, ok := claims["mobile"].(string)
	if !ok {
		return "", errors.New("illegal token")
	}

	return mobile, nil
}
