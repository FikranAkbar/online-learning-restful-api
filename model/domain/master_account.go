package domain

import (
	"online-learning-restful-api/core"
)

type MasterAccount struct {
	core.DomainModel
	Email    string `json:"email" gorm:"not null" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"required"`
}
