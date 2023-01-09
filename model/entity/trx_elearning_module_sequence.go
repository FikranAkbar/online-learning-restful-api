package entity

import (
	"gorm.io/gorm"
)

type TrxElearningModuleSequence struct {
	gorm.Model     `gorm:"embedded"`
	ModuleId       uint `gorm:"column:module_id;not null"`
	SequenceNumber uint `gorm:"column:sequence_number;type:int(10);not null"`
}

func (TrxElearningModuleSequence) TableName() string {
	return "trx_elearning_module_sequence"
}
