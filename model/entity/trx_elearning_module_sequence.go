package entity

import "online-learning-restful-api/core"

type TrxElearningModuleSequence struct {
	core.EntityModel `gorm:"embedded"`
	ModuleId         uint `gorm:"not null"`
	SequenceNumber   uint `gorm:"type:int(10);not null"`
}

func (TrxElearningModuleSequence) TableName() string {
	return "trx_elearning_module_sequence"
}
