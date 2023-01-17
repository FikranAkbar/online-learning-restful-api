package entity

import (
	"time"
)

type TrxUserVideoProgression struct {
	ID          uint `gorm:"column:id;primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	VideoId     uint        `gorm:"column:video_id;not null"`
	Video       MasterVideo `gorm:"foreignKey:VideoId;joinForeignKey:ID"`
	UserId      uint        `gorm:"column:user_id;not null"`
	User        MasterUser  `gorm:"foreignKey:UserId;joinForeignKey:ID"`
	Progression float32     `gorm:"column:progression"`
	IsComplete  bool        `gorm:"column:is_complete;default:false"`
}

func (TrxUserVideoProgression) TableName() string {
	return "trx_user_video_progression"
}
