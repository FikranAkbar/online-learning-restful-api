package entity

import "online-learning-restful-api/core"

type MasterPaymentChannel struct {
	core.EntityModel `gorm:"embedded"`
	PaymentChannel   string `gorm:"type:varchar(100);not null"`
}

func (MasterPaymentChannel) TableName() string {
	return "master_payment_channel"
}
