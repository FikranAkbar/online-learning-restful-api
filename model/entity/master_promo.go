package entity

import (
	"gorm.io/gorm"
)

type MasterPromo struct {
	gorm.Model `gorm:"embedded"`
	PromoName  string  `gorm:"column:promo_name;type:varchar(200);not null"`
	Discount   float32 `gorm:"column:discount;not null"`
}

func (MasterPromo) TableName() string {
	return "master_promo"
}
