package domain

import "online-learning-restful-api/core"

type MasterWebinarSession struct {
	core.DomainModel
	CourseId     uint   `gorm:"not null"`
	Title        string `gorm:"type:varchar(200)"`
	Desc         string
	Cover        string
	ZoomLink     string
	ScheduleDay  uint `gorm:"not null"`
	ScheduleDate string
	TimeStart    string
	TimeFinish   string
	IsPublished  bool `gorm:"default:false"`
}
