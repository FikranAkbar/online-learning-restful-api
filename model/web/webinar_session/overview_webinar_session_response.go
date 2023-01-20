package webinar_session

type OverviewWebinarSessionResponse struct {
	CourseId         uint   `json:"course_id" mapstructure:"course_id"`
	WebinarSessionId uint   `json:"webinar_session_id" mapstructure:"webinar_session_id"`
	Title            string `json:"title" mapstructure:"title"`
	Cover            string `json:"cover" mapstructure:"cover"`
	ScheduleDay      string `json:"schedule_day" mapstructure:"schedule_day"`
	ScheduleDate     string `json:"schedule_date" mapstructure:"schedule_date"`
	TimeStart        string `json:"time_start" mapstructure:"time_start"`
	TimeFinish       string `json:"time_finish" mapstructure:"time_finish"`
}
