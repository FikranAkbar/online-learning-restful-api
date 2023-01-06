package domain

import "online-learning-restful-api/core"

type MasterPaymentChannel struct {
	core.DomainModel
	PaymentChannel string `gorm:"type:varchar(100)"`
}

func (MasterPaymentChannel) TableName() string {
	return "master_payment_channel"
}
