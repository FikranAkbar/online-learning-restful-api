package entity

import (
	"database/sql"
	"time"
)

type MasterWebinarSession struct {
	ID           uint `gorm:"column:id;primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CourseId     uint                      `gorm:"column:course_id;not null"`
	MasterCourse MasterCourse              `gorm:"foreignKey:CourseId;joinForeignKey:ID"`
	Title        string                    `gorm:"column:title;type:varchar(200);not null"`
	Desc         sql.NullString            `gorm:"column:desc"`
	Cover        string                    `gorm:"column:cover;type:varchar(500);default:https://images.unsplash.com/photo-1608231975456-2f2d9fb1b49b?crop=entropy&cs=tinysrgb&fit=crop&fm=jpg&h=800&ixid=MnwxfDB8MXxyYW5kb218MHx8fHx8fHx8MTY0OTczNzM3MA&ixlib=rb-1.2.1&q=80&utm_campaign=api_test-credit&utm_medium=referral&utm_source=unsplash_source&w=1200"`
	ZoomLink     string                    `gorm:"column:zoom_link;type:varchar(500);default:https://us04web.zoom.us/j/75054156086?pwd=>LeKLsbETWFwo3-BgICsjRpUo7MFEuV.1"`
	ScheduleDay  uint                      `gorm:"column:schedule_day;"`
	Day          MasterDay                 `gorm:"foreignKey:ScheduleDay;joinForeignKey:ID"`
	ScheduleDate string                    `gorm:"column:schedule_date"`
	TimeStart    string                    `gorm:"column:time_start"`
	TimeFinish   string                    `gorm:"column:time_finish"`
	IsPublished  bool                      `gorm:"column:is_published;default:false"`
	Sequence     TrxWebinarSessionSequence `gorm:"foreignKey:WebinarSessionId;joinForeignKey:ID"`
}

func (MasterWebinarSession) TableName() string {
	return "master_webinar_session"
}
