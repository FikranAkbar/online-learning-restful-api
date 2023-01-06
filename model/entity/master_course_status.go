package entity

import "online-learning-restful-api/core"

type MasterCourseStatus struct {
	core.EntityModel
	Name string `gorm:"type:varchar(100)"`
}

func (MasterCourseStatus) TableName() string {
	return "master_course_status"
}
