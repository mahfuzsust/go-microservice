package services

import (
	"auth-service/entity"
	"auth-service/repository"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Authenticator struct {
	repository *repository.UserRepository
}

func (a *Authenticator) GenerateToken(username, password string) (string, error) {
	isAuthenticated := a.repository.IsAuthenticated(username, password)

	if !isAuthenticated {
		return "", errors.New("User not authenticated")
	}

	generator := GetGenerator("jwt")

	tokenId, _ := uuid.NewUUID()

	jwtToken, err := generator.CreateToken(username, tokenId, time.Hour)

	if err != nil {
		return "", errors.New("JWT generate error")
	}
	return jwtToken, nil
}

func (a *Authenticator) VerifyToken(tokenEntity entity.TokenEntity) (*entity.Payload, error) {
	generator := GetGenerator(strings.ToLower(tokenEntity.Type))

	payload, err := generator.VerfifyToken(tokenEntity.Token)

	if err != nil {
		return nil, errors.New("JWT verification error")
	}
	return payload, nil
}

func New() *Authenticator {
	service := &Authenticator{}
	service.repository = repository.New()
	return service
}
