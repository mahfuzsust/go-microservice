package services

import (
	"auth-service/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTGenerator struct {
	secretKey string
}

func (g *JWTGenerator) CreateToken(username string, tokenId uuid.UUID, duration time.Duration) (string, error) {
	payload := entity.NewPayload(username, tokenId, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(g.secretKey))
}

func (g *JWTGenerator) VerfifyToken(token string) (*entity.Payload, error) {
	return nil, nil
}
