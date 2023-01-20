package domain

type CourseReview struct {
	Id         uint
	CourseId   uint
	UserId     uint
	UserName   string
	ReviewDesc string
	ReviewRate float32
}
