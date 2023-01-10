package entity

import (
	"gorm.io/gorm"
)

type MasterVideo struct {
	gorm.Model  `gorm:"embedded"`
	CourseId    uint   `gorm:"column:course_id;not null"`
	ModuleId    uint   `gorm:"column:module_id;not null"`
	VideoName   string `gorm:"column:video_name;type:varchar(200);not null;unique_index"`
	VideoUrl    string `gorm:"column:video_url;type:varchar(500);default:http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"`
	IsPublished bool   `gorm:"column:is_published;default:false"`
}

func (MasterVideo) TableName() string {
	return "master_video"
}
