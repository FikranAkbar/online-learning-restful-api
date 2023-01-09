package entity

import (
	"gorm.io/gorm"
	"time"
)

type MasterCity struct {
	ID             *uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	CityName       string         `gorm:"column:city_name;type:varchar(100);not null"`
	ProvinceId     uint           `gorm:"column:province_id"`
	MasterProvince MasterProvince `gorm:"foreignKey:ProvinceId;references:ID"`
}

func (MasterCity) TableName() string {
	return "master_city"
}
