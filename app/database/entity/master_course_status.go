package entity

import (
	"time"
)

type MasterCourseStatus struct {
	ID        uint `gorm:"column:id;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"column:name;type:varchar(100);not null"`
}

func (MasterCourseStatus) TableName() string {
	return "master_course_status"
}
