package repository

import (
	"auth-service/models"
	"auth-service/utils"

	"gorm.io/gorm"
)

type UserRepository struct {
	connection *gorm.DB
}

func (r *UserRepository) IsAuthenticated(username, password string) bool {
	user := &models.UserAuth{}
	result := r.connection.First(&user, "username = ?", username)

	if result.RowsAffected == 0 {
		return false
	}

	return utils.CheckPasswordHash(password, user.Password)

}

func New() *UserRepository {
	repo := &UserRepository{}
	repo.connection = GetDBConnection().Statement.DB
	return repo
}
