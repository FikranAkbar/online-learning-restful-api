package entity

import (
	"time"
)

type MasterLikeableType struct {
	ID           uint `gorm:"column:id;primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LikeableName string `gorm:"column:likeable_name;type:varchar(50);not null"`
}

func (MasterLikeableType) TableName() string {
	return "master_likeable_type"
}
