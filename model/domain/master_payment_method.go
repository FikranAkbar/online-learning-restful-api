package domain

import "online-learning-restful-api/core"

type MasterPaymentMethod struct {
	core.DomainModel
	Name string `gorm:"type:varchar(40)"`
}
