package entity

import (
	"online-learning-restful-api/core"
)

type MasterAccount struct {
	core.EntityModel  `gorm:"embedded"`
	Email             string `gorm:"type:varchar(300);not null;unique_index"`
	Password          string `gorm:"type:varchar(300);not null"`
	Role              uint
	MasterUserType    MasterUserType      `gorm:"foreignKey:Role"`
	TrxForgotPassword []TrxForgotPassword `gorm:"foreignKey:AccountId"`
}

func (MasterAccount) TableName() string {
	return "master_account"
}
