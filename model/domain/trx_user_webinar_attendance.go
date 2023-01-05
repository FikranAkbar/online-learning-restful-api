package domain

import (
	"online-learning-restful-api/core"
	"time"
)

type TrxUserWebinarAttendance struct {
	core.DomainModel
	WebinarSessionId uint `gorm:"not null"`
	UserId           uint `gorm:"not null"`
	FirstJoinTime    time.Time
	LastLeaveTime    time.Time
	JoinDuration     uint
}
