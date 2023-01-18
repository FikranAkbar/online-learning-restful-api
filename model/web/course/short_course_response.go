package course

type ShortCourseResponse struct {
	Id          uint    `json:"id,omitempty"`
	Name        string  `json:"course_name,omitempty" mapstructure:"course_name"`
	PhotoUrl    string  `json:"photo_url,omitempty" mapstructure:"photo_url"`
	AverageRate float32 `json:"average_rate,omitempty" mapstructure:"average_rate"`
	ExpertName  string  `json:"expert_name,omitempty" mapstructure:"expert_name"`
}
