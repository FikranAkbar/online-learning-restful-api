package domain

type WebinarSession struct {
	Id             uint
	CourseId       uint
	Title          string
	Desc           string
	Cover          string
	ZoomLink       string
	ScheduleDay    string
	ScheduleDate   string
	TimeStart      string
	TimeFinish     string
	IsPublished    bool
	SequenceNumber uint
}
