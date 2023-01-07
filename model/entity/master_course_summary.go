package entity

import (
	"gorm.io/gorm"
)

type MasterCourseSummary struct {
	gorm.Model   `gorm:"embedded"`
	CourseId     uint         `gorm:"column:course_id;not null"`
	MasterCourse MasterCourse `gorm:"foreignKey:CourseId;references:ID"`
	DocURL       string       `gorm:"column:doc_url;not null"`
}

func (MasterCourseSummary) TableName() string {
	return "master_course_summary"
}
