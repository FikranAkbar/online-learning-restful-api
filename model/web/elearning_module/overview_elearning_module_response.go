package elearning_module

type OverviewElearningModuleResponse struct {
	CourseId uint   `json:"course_id" mapstructure:"course_id"`
	ModuleId uint   `json:"module_id" mapstructure:"module_id"`
	Title    string `json:"title" mapstructure:"title"`
	Overview string `json:"overview" mapstructure:"overview"`
}
