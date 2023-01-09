package entity

import (
	"database/sql"
	"gorm.io/gorm"
)

type MasterCourseSummary struct {
	gorm.Model   `gorm:"embedded"`
	CourseId     uint           `gorm:"column:course_id;not null"`
	MasterCourse MasterCourse   `gorm:"foreignKey:CourseId;references:ID"`
	DocURL       sql.NullString `gorm:"column:doc_url"`
}

func (MasterCourseSummary) TableName() string {
	return "master_course_summary"
}
