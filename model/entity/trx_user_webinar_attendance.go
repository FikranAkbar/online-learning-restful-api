package entity

import (
	"gorm.io/gorm"
	"time"
)

type TrxUserWebinarAttendance struct {
	gorm.Model       `gorm:"embedded"`
	WebinarSessionId uint      `gorm:"column:webinar_session_id;not null"`
	UserId           uint      `gorm:"column:user_id;not null"`
	FirstJoinTime    time.Time `gorm:"column:first_join_time"`
	LastLeaveTime    time.Time `gorm:"column:last_leave_time"`
	JoinDuration     uint      `gorm:"column:join_duration"`
}

func (TrxUserWebinarAttendance) TableName() string {
	return "trx_user_webinar_attendance"
}
