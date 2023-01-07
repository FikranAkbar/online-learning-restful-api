package entity

import "online-learning-restful-api/core"

type TrxElearningModuleSequence struct {
	core.EntityModel      `gorm:"embedded"`
	ModuleId              uint                  `gorm:"column:module_id;not null"`
	MasterElearningModule MasterElearningModule `gorm:"foreignKey:ModuleId"`
	SequenceNumber        uint                  `gorm:"column:sequence_number;type:int(10);not null"`
}

func (TrxElearningModuleSequence) TableName() string {
	return "trx_elearning_module_sequence"
}
