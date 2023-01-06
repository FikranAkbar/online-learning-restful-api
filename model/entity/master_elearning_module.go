package entity

import "online-learning-restful-api/core"

type MasterElearningModule struct {
	core.EntityModel
	CourseId          uint   `gorm:"not null"`
	ModuleTitle       string `gorm:"type:varchar(100)"`
	ModuleOverview    string
	ModuleDescription string
	IsPublished       bool `gorm:"default:false"`
}

func (MasterElearningModule) TableName() string {
	return "master_elearning_module"
}
