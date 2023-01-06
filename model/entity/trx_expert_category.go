package entity

import "online-learning-restful-api/core"

type TrxExpertCategory struct {
	core.EntityModel `gorm:"embedded"`
	ExpertId         uint `gorm:"not null"`
	CategoryId       uint `gorm:"not null"`
}

func (TrxExpertCategory) TableName() string {
	return "trx_expert_category"
}
