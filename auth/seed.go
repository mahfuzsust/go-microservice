package main

import (
	"auth-service/models"
	"auth-service/repository"
	"auth-service/utils"

	"github.com/google/uuid"
)

func seed() {
	conn := repository.GetDBConnection()

	conn.Statement.DB.AutoMigrate(&models.UserAuth{})

	pass, _ := utils.HashPassword("Hello world")
	userId, _ := uuid.NewUUID()
	// // Create
	conn.Statement.DB.Create(&models.UserAuth{
		Base:     models.Base{},
		Username: "test",
		Password: pass,
		UserId:   userId,
	})
}
