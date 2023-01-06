package entity

import "online-learning-restful-api/core"

type TrxUserVideoProgression struct {
	core.EntityModel `gorm:"embedded"`
	VideoId          uint           `gorm:"not null"`
	MasterVideo      MasterVideo    `gorm:"foreignKey:VideoId"`
	UserId           uint           `gorm:"not null"`
	MasterUser       MasterUserType `gorm:"foreignKey:UserId"`
	Progression      float32
	IsComplete       bool `gorm:"default:false"`
}

func (TrxUserVideoProgression) TableName() string {
	return "trx_user_video_progression"
}
