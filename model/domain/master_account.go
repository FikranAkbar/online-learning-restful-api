package domain

import (
	"online-learning-restful-api/core"
)

type MasterAccount struct {
	core.DomainModel
	Email    string `gorm:"type:varchar(300);not null;email;unique_index"`
	Password string `gorm:"type:varchar(300);not null"`
	Role     uint
}

func (MasterAccount) TableName() string {
	return "master_account"
}
