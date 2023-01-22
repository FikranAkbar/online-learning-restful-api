package course_service

import (
	"context"
	"online-learning-restful-api/model/web/course"
)

type CourseSummaryService interface {
	GetCourseSummary(ctx context.Context, courseId uint) course.SummaryResponse
}
