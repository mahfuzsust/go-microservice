package services

import (
	"auth-service/repository"
	"errors"

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

	jwtToken, err := generator.CreateToken(username, tokenId, 30)

	if err != nil {
		return "", errors.New("JWT generate error")
	}
	return jwtToken, nil
}

func New() *Authenticator {
	service := &Authenticator{}
	service.repository = repository.New()
	return service
}
