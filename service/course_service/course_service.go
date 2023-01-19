package course_service

import (
	"context"
	"online-learning-restful-api/model/web/course"
)

type CourseService interface {
	GetCoursesByKeyword(ctx context.Context, keyword string) []course.ShortCourseResponse
}
