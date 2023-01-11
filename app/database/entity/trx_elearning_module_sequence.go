package entity

import (
	"gorm.io/gorm"
	"time"
)

type TrxElearningModuleSequence struct {
	ModuleId    uint `gorm:"column:module_id;primaryKey"`
	OrderNumber uint `gorm:"column:order_number;primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (TrxElearningModuleSequence) TableName() string {
	return "trx_elearning_module_sequence"
}
