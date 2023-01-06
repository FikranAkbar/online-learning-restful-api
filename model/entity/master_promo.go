package entity

import "online-learning-restful-api/core"

type MasterPromo struct {
	core.EntityModel `gorm:"embedded"`
	PromoName        string  `gorm:"type:varchar(200);not null"`
	Discount         float32 `gorm:"not null"`
}

func (MasterPromo) TableName() string {
	return "master_promo"
}
