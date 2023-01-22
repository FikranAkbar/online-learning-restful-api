package expert

type DetailExpertResponse struct {
	ExpertId           uint    `json:"expert_id" mapstructure:"expert_id"`
	Name               string  `json:"name" mapstructure:"name"`
	ProfileDescription string  `json:"profile_description,omitempty" mapstructure:"profile_description"`
	Photo              string  `json:"photo" mapstructure:"photo"`
	AverageRate        float32 `json:"average_rate" mapstructure:"average_rate"`
	TotalStudent       int     `json:"total_student" mapstructure:"total_student"`
	TotalCourse        int     `json:"total_course" mapstructure:"total_course"`
}
