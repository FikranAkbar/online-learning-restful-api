package entity

import (
	"time"
)

type TrxUserWebinarAttendance struct {
	ID               uint `gorm:"column:id;primarykey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	WebinarSessionId uint                 `gorm:"column:webinar_session_id;not null"`
	WebinarSession   MasterWebinarSession `gorm:"foreignKey:WebinarSessionId;joinForeignKey:ID"`
	UserId           uint                 `gorm:"column:user_id;not null"`
	User             MasterUser           `gorm:"foreignKey:UserId;joinForeignKey:ID"`
	FirstJoinTime    time.Time            `gorm:"column:first_join_time"`
	LastLeaveTime    time.Time            `gorm:"column:last_leave_time"`
	JoinDuration     uint                 `gorm:"column:join_duration"`
}

func (TrxUserWebinarAttendance) TableName() string {
	return "trx_user_webinar_attendance"
}
