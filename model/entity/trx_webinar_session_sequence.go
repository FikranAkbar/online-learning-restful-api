package entity

import "online-learning-restful-api/core"

type TrxWebinarSessionSequence struct {
	core.EntityModel     `gorm:"embedded"`
	WebinarSessionId     uint                 `gorm:"column:webinar_session_id;not null"`
	MasterWebinarSession MasterWebinarSession `gorm:"foreignKey:WebinarSessionId"`
	SequenceNumber       uint                 `gorm:"column:sequence_number;not null"`
}

func (TrxWebinarSessionSequence) TableName() string {
	return "trx_webinar_session_sequence"
}
