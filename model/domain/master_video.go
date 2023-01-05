package domain

import "online-learning-restful-api/core"

type MasterVideo struct {
	core.DomainModel
	CourseId    uint   `gorm:"not null"`
	ModuleId    uint   `gorm:"not null"`
	VideoName   string `gorm:"type:varchar(100)"`
	VideoUrl    string
	IsPublished bool `gorm:"default:false"`
}
