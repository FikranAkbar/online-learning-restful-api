package entity

import (
	"time"
)

type MasterDay struct {
	ID        uint `gorm:"column:id;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"column:name;type:varchar(20);not null"`
}

func (MasterDay) TableName() string {
	return "master_day"
}
