package domain

import "online-learning-restful-api/core"

type MasterCourseStatus struct {
	core.DomainModel
	Status string `gorm:"type:varchar(100)"`
}
