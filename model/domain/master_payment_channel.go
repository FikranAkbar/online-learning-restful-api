package domain

import "online-learning-restful-api/core"

type MasterPaymentChannel struct {
	core.DomainModel
	PaymentChannel string `gorm:"type:varchar(100)"`
}
