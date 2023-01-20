package webinar_session_service

import (
	"context"
	"online-learning-restful-api/model/web/webinar_session"
)

type WebinarSessionService interface {
	GetOverviewWebinarSessionsByCourseId(ctx context.Context, courseId uint) []webinar_session.OverviewWebinarSessionResponse
	GetDetailWebinarSessionByWebinarSessionId(ctx context.Context, courseId uint, webinarSessionId uint) webinar_session.DetailWebinarSessionResponse
}
