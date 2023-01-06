package entity

import "online-learning-restful-api/core"

type MasterPromo struct {
	core.EntityModel
	PromoName string `gorm:"type:varchar(200)"`
	Discount  float32
}

func (MasterPromo) TableName() string {
	return "master_promo"
}
