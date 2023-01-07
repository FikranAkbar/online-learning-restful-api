package entity

import (
	"gorm.io/gorm"
)

type MasterPaymentStatus struct {
	gorm.Model `gorm:"embedded"`
	Name       string `gorm:"column:name;type:varchar(100);not null"`
}

func (MasterPaymentStatus) TableName() string {
	return "master_payment_status"
}
