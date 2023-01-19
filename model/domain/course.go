package domain

type Course struct {
	Id                 uint
	ExpertId           uint
	ExpertName         string
	StatusCourse       string
	Name               string
	Description        string
	PhotoUrl           string
	AverageRate        float32
	Price              uint
	TotalRate          uint
	TotalDuration      float32
	CurrentParticipant uint
	MaximumParticipant uint
	AlreadyOwned       bool
	ModulesCount       int
	WebinarsCount      int
}
