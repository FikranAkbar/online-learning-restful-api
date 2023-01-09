package entity

import (
	"gorm.io/gorm"
)

type MasterAccount struct {
	gorm.Model        `gorm:"embedded"`
	Email             string              `gorm:"column:email;type:varchar(300);not null;unique_index"`
	Password          string              `gorm:"column:password;type:varchar(300);not null"`
	Role              uint                `gorm:"column:role"`
	MasterUserType    MasterUserType      `gorm:"foreignKey:Role"`
	TrxForgotPassword []TrxForgotPassword `gorm:"foreignKey:AccountId;references:ID"`
}

func (MasterAccount) TableName() string {
	return "master_account"
}