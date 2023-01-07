package entity

import (
	"gorm.io/gorm"
)

type MasterPaymentChannel struct {
	gorm.Model     `gorm:"embedded"`
	PaymentChannel string `gorm:"column:payment_channel;type:varchar(100);not null"`
}

func (MasterPaymentChannel) TableName() string {
	return "master_payment_channel"
}
