package entity

import "online-learning-restful-api/core"

type MasterCourseStatus struct {
	core.EntityModel `gorm:"embedded"`
	Name             string `gorm:"type:varchar(100);not null"`
}

func (MasterCourseStatus) TableName() string {
	return "master_course_status"
}
