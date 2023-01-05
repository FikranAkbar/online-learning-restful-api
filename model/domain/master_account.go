package domain

import (
	"online-learning-restful-api/core"
)

type MasterAccount struct {
	core.DomainModel
	Email    string `gorm:"type:varchar(255);not null;email;unique_index"`
	Password string `gorm:"type:varchar(20);not null"`
}
