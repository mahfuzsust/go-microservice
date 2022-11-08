package main

import (
	"auth-service/models"
	"auth-service/repository"
)

func seed() {
	conn := repository.GetDBConnection()

	conn.Statement.DB.AutoMigrate(&models.UserAuth{})
}
