package entity

import "online-learning-restful-api/core"

type TrxForgotPassword struct {
	core.EntityModel `gorm:"embedded"`
	AccountId        uint   `gorm:"not null"`
	Token            string `gorm:"not null"`
	IsConsumed       bool   `gorm:"default:false"`
}

func (TrxForgotPassword) TableName() string {
	return "trx_forgot_password"
}
