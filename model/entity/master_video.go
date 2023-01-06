package entity

import "online-learning-restful-api/core"

type MasterVideo struct {
	core.EntityModel `gorm:"embedded"`
	CourseId         uint   `gorm:"not null"`
	ModuleId         uint   `gorm:"not null"`
	VideoName        string `gorm:"type:varchar(100);not null;unique_index"`
	VideoUrl         string `gorm:"not null"`
	IsPublished      bool   `gorm:"default:false"`
}

func (MasterVideo) TableName() string {
	return "master_video"
}
