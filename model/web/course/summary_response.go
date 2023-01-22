package course

type SummaryResponse struct {
	Id       uint   `json:"id" mapstructure:"id"`
	CourseId uint   `json:"course_id" mapstructure:"course_id"`
	DocURL   string `json:"doc_url" mapstructure:"doc_url"`
}
