package entity

import (
	"gorm.io/gorm"
)

type TrxElearningModuleSequence struct {
	gorm.Model  `gorm:"embedded"`
	ModuleId    uint `gorm:"column:module_id;not null"`
	OrderNumber uint `gorm:"column:order_number;type:int(10);not null"`
}

func (TrxElearningModuleSequence) TableName() string {
	return "trx_elearning_module_sequence"
}
