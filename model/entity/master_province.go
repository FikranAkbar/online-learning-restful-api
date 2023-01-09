package entity

import (
	"gorm.io/gorm"
	"time"
)

type MasterProvince struct {
	ID           *uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	ProvinceName string         `gorm:"column:province_name;type:varchar(100);not null"`
}

func (MasterProvince) TableName() string {
	return "master_province"
}
