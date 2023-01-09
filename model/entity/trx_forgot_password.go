package entity

import (
	"gorm.io/gorm"
)

type TrxForgotPassword struct {
	gorm.Model `gorm:"embedded"`
	AccountId  uint   `gorm:"column:account_id;not null"`
	Token      string `gorm:"column:token;not null"`
	IsConsumed bool   `gorm:"column:is_consumed;default:false"`
}

func (TrxForgotPassword) TableName() string {
	return "trx_forgot_password"
}
