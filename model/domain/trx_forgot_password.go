package domain

import "online-learning-restful-api/core"

type TrxForgotPassword struct {
	core.DomainModel
	AccountId  uint `gorm:"not null"`
	Token      string
	IsConsumed bool `gorm:"default:false"`
}

func (TrxForgotPassword) TableName() string {
	return "trx_forgot_password"
}
