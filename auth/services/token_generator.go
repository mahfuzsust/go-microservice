package services

import (
	"auth-service/entity"
	"time"

	"github.com/google/uuid"
)

type TokenGenerator interface {
	CreateToken(username string, tokenId uuid.UUID, duration time.Duration) (string, error)
	VerfifyToken(token string) (*entity.Payload, error)
}

func GetGenerator(generatorType string) TokenGenerator {
	switch generatorType {
	case "jwt":
		return &JWTGenerator{
			secretKey: "fklsjfoiewkewnfwklefjowieurwjefwklefklfjwkejfoiwerewoir",
		}
	default:
		return &JWTGenerator{
			secretKey: "fklsjfoiewkewnfwklefjowieurwjefwklefklfjwkejfoiwerewoir",
		}
	}
}
