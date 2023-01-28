package course

type OrderCourseResponse struct {
	UserId   uint `json:"user_id" mapstructure:"user_id"`
	CourseId uint `json:"course_id" mapstructure:"course_id"`
}
