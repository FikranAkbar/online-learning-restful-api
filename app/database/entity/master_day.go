package entity

import (
	"gorm.io/gorm"
)

type MasterDay struct {
	gorm.Model `gorm:"embedded"`
	Name       string `gorm:"column:name;type:varchar(20);not null"`
}

func (MasterDay) TableName() string {
	return "master_day"
}
