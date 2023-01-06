package entity

import (
	"database/sql"
	"online-learning-restful-api/core"
	"time"
)

type MasterWebinarSession struct {
	core.EntityModel `gorm:"embedded"`
	CourseId         uint   `gorm:"not null"`
	Title            string `gorm:"type:varchar(200);not null"`
	Desc             sql.NullString
	Cover            sql.NullString
	ZoomLink         sql.NullString
	ScheduleDay      uint `gorm:"not null"`
	ScheduleDate     time.Time
	TimeStart        time.Time
	TimeFinish       time.Time
	IsPublished      bool `gorm:"default:false"`
}

func (MasterWebinarSession) TableName() string {
	return "master_webinar_session"
}
