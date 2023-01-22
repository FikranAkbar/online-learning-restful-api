package entity

import (
	"time"
)

type TrxUserVideoProgression struct {
	ID          uint `gorm:"column:id;primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	VideoId     uint       `gorm:"column:video_id;not null"`
	UserId      uint       `gorm:"column:user_id;not null"`
	User        MasterUser `gorm:"foreignKey:UserId;joinForeignKey:ID"`
	Progression float32    `gorm:"column:progression"`
	IsComplete  bool       `gorm:"column:is_complete;type:boolean"`
}

func (TrxUserVideoProgression) TableName() string {
	return "trx_user_video_progression"
}
