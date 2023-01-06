package domain

import "online-learning-restful-api/core"

type MasterUserType struct {
	core.DomainModel
	UserType string `gorm:"type:varchar(20)"`
}

func (MasterUserType) TableName() string {
	return "master_user_type"
}
