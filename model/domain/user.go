package domain

import (
	"time"
)

type User struct {
	Id           uint
	Name         string
	BirthDate    time.Time
	Gender       string
	Phone        string
	PhotoURL     string
	ProvinceId   uint
	ProvinceName string
}
