package entity

import (
	"online-learning-restful-api/core"
	"time"
)

type TrxUserWebinarAttendance struct {
	core.EntityModel     `gorm:"embedded"`
	WebinarSessionId     uint                 `gorm:"column:webinar_session_id;not null"`
	MasterWebinarSession MasterWebinarSession `gorm:"foreignKey:WebinarSessionId"`
	UserId               uint                 `gorm:"column:user_id;not null"`
	MasterUser           MasterUser           `gorm:"foreignKey:UserId"`
	FirstJoinTime        time.Time            `gorm:"column:first_join_time"`
	LastLeaveTime        time.Time            `gorm:"column:last_leave_time"`
	JoinDuration         uint                 `gorm:"column:join_duration"`
}

func (TrxUserWebinarAttendance) TableName() string {
	return "trx_user_webinar_attendance"
}
