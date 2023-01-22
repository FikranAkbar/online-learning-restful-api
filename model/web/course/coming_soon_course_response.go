package course

type ComingSoonCourseResponse struct {
	Id         uint   `json:"id" mapstructure:"id"`
	CourseName string `json:"course_name" mapstructure:"course_name"`
	PhotoUrl   string `json:"photo_url" mapstructure:"photo_url"`
}
