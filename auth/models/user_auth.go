package models

import "github.com/google/uuid"

type UserAuth struct {
	Base
	Username string    `gorm:"column:username;not null;index;"`
	Password string    `gorm:"column:password;not null;"`
	UserId   uuid.UUID `gorm:"column:user_id;not null;uniqueIndex;"`
}
