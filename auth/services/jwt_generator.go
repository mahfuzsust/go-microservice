package services

import (
	"auth-service/entity"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTGenerator struct {
	secretKey string
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

func (g *JWTGenerator) CreateToken(username string, tokenId uuid.UUID, duration time.Duration) (string, error) {
	payload := entity.NewPayload(username, tokenId, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(g.secretKey))
}

func (g *JWTGenerator) VerfifyToken(token string) (*entity.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(g.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &entity.Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*entity.Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
