package entity

import "online-learning-restful-api/core"

type MasterDay struct {
	core.EntityModel `gorm:"embedded"`
	DayName          string `gorm:"type:varchar(20);not null"`
}

func (MasterDay) TableName() string {
	return "master_day"
}
