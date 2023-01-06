package entity

import "online-learning-restful-api/core"

type MasterCategory struct {
	core.EntityModel `gorm:"embedded"`
	CategoryName     string `gorm:"type:varchar(100);not null"`
}

func (MasterCategory) TableName() string {
	return "master_category"
}
