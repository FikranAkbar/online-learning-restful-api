package entity

import (
	"time"
)

type TrxWebinarSessionSequence struct {
	WebinarSessionId uint `gorm:"column:webinar_session_id;primaryKey"`
	OrderNumber      uint `gorm:"column:order_number;primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (TrxWebinarSessionSequence) TableName() string {
	return "trx_webinar_session_sequence"
}
