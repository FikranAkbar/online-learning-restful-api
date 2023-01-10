package entity

import (
	"gorm.io/gorm"
	"time"
)

type MasterPromo struct {
	ID        *uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	PromoName string         `gorm:"column:promo_name;type:varchar(200);not null"`
	Discount  float32        `gorm:"column:discount;not null"`
}

func (MasterPromo) TableName() string {
	return "master_promo"
}
