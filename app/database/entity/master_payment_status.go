package entity

import (
	"time"
)

type MasterPaymentStatus struct {
	ID        uint `gorm:"column:id;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"column:name;type:varchar(100);not null"`
}

func (MasterPaymentStatus) TableName() string {
	return "master_payment_status"
}
