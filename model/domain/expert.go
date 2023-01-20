package domain

import "time"

type Expert struct {
	Id           uint
	Name         string
	Profession   string
	Phone        string
	Address      string
	Gender       string
	Photo        string
	BirthDate    time.Time
	AverageRate  float32
	TotalStudent int
	ReviewsCount int
}
