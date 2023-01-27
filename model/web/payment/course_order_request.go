package payment

type CourseOrderRequest struct {
	CourseIds []uint `json:"course_ids" mapstructure:"course_ids"`
}
