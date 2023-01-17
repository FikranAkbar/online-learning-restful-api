package entity

import (
	"time"
)

type MasterCourseSummary struct {
	ID           uint `gorm:"column:id;primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CourseId     uint         `gorm:"column:course_id;not null"`
	MasterCourse MasterCourse `gorm:"foreignKey:CourseId;references:ID"`
	DocURL       string       `gorm:"column:doc_url;default:http://www.africau.edu/images/default/sample.pdf"`
}

func (MasterCourseSummary) TableName() string {
	return "master_course_summary"
}
