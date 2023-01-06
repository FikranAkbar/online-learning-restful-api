package entity

import "online-learning-restful-api/core"

type MasterUserType struct {
	core.EntityModel `gorm:"embedded"`
	UserType         string `gorm:"type:varchar(20);unique_index"`
}

func (MasterUserType) TableName() string {
	return "master_user_type"
}
