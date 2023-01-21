package domain

type Video struct {
	Id                   uint
	CourseId             uint
	ModuleId             uint
	VideoName            string
	VideoUrl             string
	LastVideoProgression float32
	IsPublished          bool
}
