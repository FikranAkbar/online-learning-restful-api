package entity

import (
	"gorm.io/gorm"
	"time"
)

type TrxExpertCategory struct {
	ExpertId   uint `gorm:"primaryKey"`
	CategoryId uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (TrxExpertCategory) TableName() string {
	return "trx_expert_category"
}
