package entity

import (
	"online-learning-restful-api/core"
)

type MasterAccount struct {
	core.EntityModel
	Email          string `gorm:"type:varchar(300);not null;unique_index"`
	Password       string `gorm:"type:varchar(300);not null"`
	Role           uint
	MasterUserType MasterUserType `gorm:"foreignKey:Role;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (MasterAccount) TableName() string {
	return "master_account"
}
