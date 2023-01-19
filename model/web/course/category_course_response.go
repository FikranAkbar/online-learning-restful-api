package course

type CategoryCourseResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"category_name" mapstructure:"category_name"`
}
