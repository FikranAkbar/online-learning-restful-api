package domain

import "online-learning-restful-api/core"

type MasterUserType struct {
	core.DomainModel
	UserType string `gorm:"type:varchar(20)"`
}
