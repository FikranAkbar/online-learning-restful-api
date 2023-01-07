package entity

import "online-learning-restful-api/core"

type MasterVideo struct {
	core.EntityModel      `gorm:"embedded"`
	CourseId              uint                  `gorm:"column:course_id;not null"`
	MasterCourse          MasterCourse          `gorm:"foreignKey:CourseId"`
	ModuleId              uint                  `gorm:"column:module_id;not null"`
	MasterElearningModule MasterElearningModule `gorm:"foreignKey:ModuleId"`
	VideoName             string                `gorm:"column:video_name;type:varchar(100);not null;unique_index"`
	VideoUrl              string                `gorm:"column:video_url;not null"`
	IsPublished           bool                  `gorm:"column:is_published;default:false"`
}

func (MasterVideo) TableName() string {
	return "master_video"
}
