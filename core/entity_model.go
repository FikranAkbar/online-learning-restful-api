package core

import "time"

type EntityModel struct {
	ID        uint       `gorm:"column:id;primaryKey"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
