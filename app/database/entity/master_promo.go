package entity

import (
	"time"
)

type MasterPromo struct {
	ID        *uint `gorm:"column:id;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	PromoName string  `gorm:"column:promo_name;type:varchar(200);not null"`
	Discount  float32 `gorm:"column:discount;not null"`
}

func (MasterPromo) TableName() string {
	return "master_promo"
}
