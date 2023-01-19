package course_service

import (
	"context"
	"online-learning-restful-api/model/web/course"
)

type CoursePopularService interface {
	GetPopularCourses(ctx context.Context) []course.ShortCourseResponse
}
