package entity

import "online-learning-restful-api/core"

type MasterPaymentStatus struct {
	core.EntityModel `gorm:"embedded"`
	Name             string `gorm:"type:varchar(100);not null"`
}

func (MasterPaymentStatus) TableName() string {
	return "master_payment_status"
}
