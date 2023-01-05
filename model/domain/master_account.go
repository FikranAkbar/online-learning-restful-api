package domain

import (
	"online-learning-restful-api/core"
)

type MasterAccount struct {
	core.DomainModel
	Email    string `gorm:"not null,email"`
	Password string `gorm:"not null"`
}
