package entity

import (
	"gorm.io/gorm"
)

type TrxUserVideoProgression struct {
	gorm.Model  `gorm:"embedded"`
	VideoId     uint    `gorm:"column:video_id;not null"`
	UserId      uint    `gorm:"column:user_id;not null"`
	Progression float32 `gorm:"column:progression"`
	IsComplete  bool    `gorm:"column:is_complete;default:false"`
}

func (TrxUserVideoProgression) TableName() string {
	return "trx_user_video_progression"
}
