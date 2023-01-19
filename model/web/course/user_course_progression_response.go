package course

import "time"

type UserCourseProgressionResponse struct {
	UserId             uint      `json:"user_id,omitempty"`
	CourseId           uint      `json:"course_id,omitempty"`
	LastUnlockedModule uint      `json:"last_unlocked_module,omitempty"`
	GraduatedAt        time.Time `json:"graduated_at,omitempty"`
}
