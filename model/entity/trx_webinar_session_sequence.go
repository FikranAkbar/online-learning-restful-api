package entity

import (
	"gorm.io/gorm"
)

type TrxWebinarSessionSequence struct {
	gorm.Model       `gorm:"embedded"`
	WebinarSessionId uint `gorm:"column:webinar_session_id;not null"`
	SequenceNumber   uint `gorm:"column:sequence_number;not null"`
}

func (TrxWebinarSessionSequence) TableName() string {
	return "trx_webinar_session_sequence"
}
