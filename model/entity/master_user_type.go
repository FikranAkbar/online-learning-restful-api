package entity

import (
	"gorm.io/gorm"
)

type MasterUserType struct {
	gorm.Model `gorm:"embedded"`
	UserType   string `gorm:"column:user_type;type:varchar(20);unique_index"`
}

func (MasterUserType) TableName() string {
	return "master_user_type"
}
