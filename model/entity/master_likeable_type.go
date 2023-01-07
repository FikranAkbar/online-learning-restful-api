package entity

import (
	"gorm.io/gorm"
)

type MasterLikeableType struct {
	gorm.Model   `gorm:"embedded"`
	LikeableName string `gorm:"column:likeable_name;type:varchar(50);not null"`
}

func (MasterLikeableType) TableName() string {
	return "master_likeable_type"
}
