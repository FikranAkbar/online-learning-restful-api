package course

type ReviewCourseResponse struct {
	CourseId   uint    `json:"course_id,omitempty" mapstructure:"course_id"`
	UserId     uint    `json:"user_id" mapstructure:"user_id"`
	UserName   string  `json:"user_name" mapstructure:"user_name"`
	ReviewDesc string  `json:"review_desc" mapstructure:"review_desc"`
	ReviewRate float32 `json:"review_rate" mapstructure:"review_rate"`
}
