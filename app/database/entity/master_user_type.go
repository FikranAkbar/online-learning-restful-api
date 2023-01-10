package entity

import (
	"gorm.io/gorm"
)

type MasterUserType struct {
	gorm.Model `gorm:"embedded"`
	Name       string `gorm:"column:name;type:varchar(20);unique_index"`
}

func (MasterUserType) TableName() string {
	return "master_user_type"
}
