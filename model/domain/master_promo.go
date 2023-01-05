package domain

import "online-learning-restful-api/core"

type MasterPromo struct {
	core.DomainModel
	PromoName string `gorm:"type:varchar(200)"`
	Discount  float32
}
