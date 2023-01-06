package entity

import "online-learning-restful-api/core"

type MasterPaymentMethod struct {
	core.EntityModel `gorm:"embedded"`
	Name             string `gorm:"type:varchar(40);not null"`
}

func (MasterPaymentMethod) TableName() string {
	return "master_payment_method"
}
