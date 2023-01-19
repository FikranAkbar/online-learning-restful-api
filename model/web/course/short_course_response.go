package course

type ShortCourseResponse struct {
	Id          uint    `json:"id"`
	Name        string  `json:"course_name" mapstructure:"course_name"`
	PhotoUrl    string  `json:"photo_url" mapstructure:"photo_url"`
	AverageRate float32 `json:"average_rate" mapstructure:"average_rate"`
	ExpertName  string  `json:"expert_name" mapstructure:"expert_name"`
}
