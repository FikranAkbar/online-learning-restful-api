package entity

import "online-learning-restful-api/core"

type TrxUserVideoProgression struct {
	core.EntityModel `gorm:"embedded"`
	VideoId          uint           `gorm:"column:video_id;not null"`
	MasterVideo      MasterVideo    `gorm:"foreignKey:VideoId"`
	UserId           uint           `gorm:"column:user_id;not null"`
	MasterUser       MasterUserType `gorm:"foreignKey:UserId"`
	Progression      float32        `gorm:"column:progression"`
	IsComplete       bool           `gorm:"column:is_complete;default:false"`
}

func (TrxUserVideoProgression) TableName() string {
	return "trx_user_video_progression"
}
