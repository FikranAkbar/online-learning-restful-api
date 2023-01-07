package entity

import (
	"gorm.io/gorm"
)

type MasterCategory struct {
	gorm.Model   `gorm:"embedded"`
	CategoryName string `gorm:"column:category_name;type:varchar(100);not null"`
}

func (MasterCategory) TableName() string {
	return "master_category"
}
