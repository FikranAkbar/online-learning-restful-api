package domain

import "online-learning-restful-api/core"

type TrxElearningModuleSequence struct {
	core.DomainModel
	ModuleId       uint `gorm:"not null"`
	SequenceNumber uint `gorm:"type:int(10)"`
}
