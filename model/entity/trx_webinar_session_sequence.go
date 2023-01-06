package entity

import "online-learning-restful-api/core"

type TrxWebinarSessionSequence struct {
	core.EntityModel
	WebinarSessionId uint `gorm:"not null"`
	SequenceNumber   uint `gorm:"not null"`
}

func (TrxWebinarSessionSequence) TableName() string {
	return "trx_webinar_session_sequence"
}
