package entity

import (
	"time"
)

type TrxElearningModuleSequence struct {
	ModuleId    uint `gorm:"column:module_id;primaryKey"`
	OrderNumber uint `gorm:"column:order_number;primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (TrxElearningModuleSequence) TableName() string {
	return "trx_elearning_module_sequence"
}
