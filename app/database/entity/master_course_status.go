package entity

import (
	"gorm.io/gorm"
)

type MasterCourseStatus struct {
	gorm.Model `gorm:"embedded"`
	Name       string `gorm:"column:name;type:varchar(100);not null"`
}

func (MasterCourseStatus) TableName() string {
	return "master_course_status"
}
