package entity

import (
	"time"
)

type MasterProvince struct {
	ID           uint `gorm:"column:id;primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProvinceName string `gorm:"column:province_name;type:varchar(100);not null"`
}

func (MasterProvince) TableName() string {
	return "master_province"
}
