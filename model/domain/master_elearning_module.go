package domain

import "online-learning-restful-api/core"

type MasterElearningModule struct {
	core.DomainModel
	CourseId          uint   `gorm:"not null"`
	ModuleTitle       string `gorm:"type:varchar(100)"`
	ModuleOverview    string
	ModuleDescription string
	IsPublished       bool `gorm:"default:false"`
}
