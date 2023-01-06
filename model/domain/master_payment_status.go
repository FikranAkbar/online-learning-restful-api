package domain

import "online-learning-restful-api/core"

type MasterPaymentStatus struct {
	core.DomainModel
	Name string `gorm:"type:varchar(100)"`
}

func (MasterPaymentStatus) TableName() string {
	return "master_payment_status"
}
