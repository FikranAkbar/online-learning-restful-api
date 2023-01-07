package entity

import (
	"gorm.io/gorm"
	"time"
)

type TrxExpertCategory struct {
	ExpertId   uint           `gorm:"column:expert_id;primaryKey"`
	CategoryId uint           `gorm:"column:category_id;primaryKey"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (TrxExpertCategory) TableName() string {
	return "trx_expert_category"
}
