package course

import "online-learning-restful-api/model/web/expert"

type DetailCourseResponse struct {
	Id                 uint                       `json:"course_id" mapstructure:"course_name"`
	Name               string                     `json:"course_name" mapstructure:"course_name"`
	Description        string                     `json:"description" mapstructure:"description"`
	PhotoUrl           string                     `json:"photo_url" mapstructure:"photo_url"`
	Contents           string                     `json:"contents" mapstructure:"contents"`
	AverageRate        float32                    `json:"average_rate" mapstructure:"average_rate"`
	CurrentParticipant uint                       `json:"current_participant" mapstructure:"current_participant"`
	MaximumParticipant uint                       `json:"maximum_participant" mapstructure:"maximum_participant"`
	AlreadyOwned       bool                       `json:"already_owned" mapstructure:"already_owned"`
	Expert             expert.ShortExpertResponse `json:"expert" mapstructure:"expert"`
}
