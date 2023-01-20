package domain

import "time"

type UserCourse struct {
	UserId               uint
	CourseId             uint
	LastUnlockedModule   uint
	TotalWebinarAttended uint
	GraduatedAt          time.Time
}
