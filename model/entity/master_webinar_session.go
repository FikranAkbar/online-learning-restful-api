package entity

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type MasterWebinarSession struct {
	gorm.Model   `gorm:"embedded"`
	CourseId     uint           `gorm:"column:course_id;not null"`
	MasterCourse MasterCourse   `gorm:"foreignKey:CourseId;joinForeignKey:ID"`
	Title        string         `gorm:"column:title;type:varchar(200);not null"`
	Desc         sql.NullString `gorm:"column:desc"`
	Cover        sql.NullString `gorm:"column:cover"`
	ZoomLink     sql.NullString `gorm:"column:zoom_link"`
	ScheduleDay  uint           `gorm:"column:schedule_day;not null"`
	ScheduleDate time.Time      `gorm:"column:schedule_date"`
	TimeStart    time.Time      `gorm:"column:time_start"`
	TimeFinish   time.Time      `gorm:"column:time_finish"`
	IsPublished  bool           `gorm:"column:is_published;default:false"`
}

func (MasterWebinarSession) TableName() string {
	return "master_webinar_session"
}
