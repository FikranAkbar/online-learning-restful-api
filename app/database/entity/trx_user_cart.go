package entity

import (
	"time"
)

type TrxUserCart struct {
	UserId    uint         `gorm:"column:user_id;primaryKey"`
	User      MasterUser   `gorm:"foreignKey:UserId;joinForeignKey:ID"`
	CourseId  uint         `gorm:"column:course_id;primaryKey"`
	Course    MasterCourse `gorm:"foreignKey:CourseId;joinForeignKey:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (TrxUserCart) TableName() string {
	return "trx_user_cart"
}
