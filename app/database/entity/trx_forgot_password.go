package entity

import (
	"time"
)

type TrxForgotPassword struct {
	ID         uint `gorm:"column:id;primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	AccountId  uint   `gorm:"column:account_id;not null"`
	Token      string `gorm:"column:token;not null"`
	IsConsumed bool   `gorm:"column:is_consumed;default:false"`
}

func (TrxForgotPassword) TableName() string {
	return "trx_forgot_password"
}
