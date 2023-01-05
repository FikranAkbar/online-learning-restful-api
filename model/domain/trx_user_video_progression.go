package domain

import "online-learning-restful-api/core"

type TrxUserVideoProgression struct {
	core.DomainModel
	VideoId     uint `gorm:"not null"`
	UserId      uint `gorm:"not null"`
	Progression float32
	IsComplete  bool `gorm:"default:false"`
}
