package course

import "time"

type UserCourseProgressionResponse struct {
	UserId             uint      `json:"user_id,omitempty" mapstructure:"user_id"`
	CourseId           uint      `json:"course_id,omitempty" mapstructure:"course_id"`
	LastUnlockedModule uint      `json:"last_unlocked_module,omitempty" mapstructure:"last_unlocked_module"`
	GraduatedAt        time.Time `json:"graduated_at,omitempty" mapstructure:"graduated_at"`
}
