package webinar_session_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/webinar_session"
	"online-learning-restful-api/repository/webinar_session_repository"
)

type WebinarSessionServiceImpl struct {
	webinar_session_repository.WebinarSessionRepository
	*gorm.DB
}

func NewWebinarSessionServiceImpl(repository webinar_session_repository.WebinarSessionRepository, DB *gorm.DB) *WebinarSessionServiceImpl {
	return &WebinarSessionServiceImpl{WebinarSessionRepository: repository, DB: DB}
}

func (service *WebinarSessionServiceImpl) GetOverviewWebinarSessionsByCourseId(ctx context.Context, courseId uint) []webinar_session.OverviewWebinarSessionResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	webinarSessions, err := service.WebinarSessionRepository.GetOverviewWebinarSessionsByCourseId(ctx, tx, courseId)
	helper.PanicIfError(err)

	var overviewWebinarSessionsResponse []webinar_session.OverviewWebinarSessionResponse
	for _, webinarSession := range webinarSessions {
		overviewWebinarSessionsResponse = append(overviewWebinarSessionsResponse, webinar_session.OverviewWebinarSessionResponse{
			CourseId:         webinarSession.CourseId,
			WebinarSessionId: webinarSession.Id,
			Title:            webinarSession.Title,
			Cover:            webinarSession.Cover,
			ScheduleDay:      webinarSession.ScheduleDay,
			ScheduleDate:     webinarSession.ScheduleDate,
			TimeStart:        webinarSession.TimeStart,
			TimeFinish:       webinarSession.TimeFinish,
		})
	}

	return overviewWebinarSessionsResponse
}
