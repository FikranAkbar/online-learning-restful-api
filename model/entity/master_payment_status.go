package entity

import "online-learning-restful-api/core"

type MasterPaymentStatus struct {
	core.EntityModel
	Name string `gorm:"type:varchar(100)"`
}

func (MasterPaymentStatus) TableName() string {
	return "master_payment_status"
}
