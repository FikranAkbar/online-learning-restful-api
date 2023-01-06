package entity

import "online-learning-restful-api/core"

type TrxUserVideoProgression struct {
	core.EntityModel `gorm:"embedded"`
	VideoId          uint `gorm:"not null"`
	UserId           uint `gorm:"not null"`
	Progression      float32
	IsComplete       bool `gorm:"default:false"`
}

func (TrxUserVideoProgression) TableName() string {
	return "trx_user_video_progression"
}
