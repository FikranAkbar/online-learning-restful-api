package entity

import (
	"gorm.io/gorm"
)

type MasterPaymentMethod struct {
	gorm.Model `gorm:"embedded"`
	Name       string `gorm:"column:name;type:varchar(40);not null"`
}

func (MasterPaymentMethod) TableName() string {
	return "master_payment_method"
}
