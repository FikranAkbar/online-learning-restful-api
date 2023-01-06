package entity

import "online-learning-restful-api/core"

type TrxExpertCategory struct {
	core.EntityModel `gorm:"embedded"`
	ExpertId         uint             `gorm:"not null"`
	MasterExpert     []MasterExpert   `gorm:"foreignKey:ExpertId"`
	CategoryId       uint             `gorm:"not null"`
	MasterCategory   []MasterCategory `gorm:"foreignKey:CategoryId"`
}

func (TrxExpertCategory) TableName() string {
	return "trx_expert_category"
}
