package domain

import "time"

type User struct {
	Id        int
	Name      string
	BirthDate time.Time
	Gender    string
}
