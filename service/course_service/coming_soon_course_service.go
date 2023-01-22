package course_service

import (
	"context"
	"online-learning-restful-api/model/web/course"
)

type ComingSoonCourseService interface {
	GetComingSoonCourses(ctx context.Context) []course.ComingSoonCourseResponse
}
