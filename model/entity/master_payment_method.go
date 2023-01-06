package entity

import "online-learning-restful-api/core"

type MasterPaymentMethod struct {
	core.EntityModel
	Name string `gorm:"type:varchar(40)"`
}

func (MasterPaymentMethod) TableName() string {
	return "master_payment_method"
}
