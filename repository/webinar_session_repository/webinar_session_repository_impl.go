package webinar_session_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/model/domain"
)

type WebinarSessionRepositoryImpl struct {
}

func NewWebinarSessionRepositoryImpl() *WebinarSessionRepositoryImpl {
	return &WebinarSessionRepositoryImpl{}
}

func (repository *WebinarSessionRepositoryImpl) GetOverviewWebinarSessionsByCourseId(ctx context.Context, db *gorm.DB, courseId uint) ([]domain.WebinarSession, error) {
	var courseEntity entity.MasterCourse
	err := db.WithContext(ctx).
		Where("id = ?", courseId).
		Preload("Webinars").
		First(&courseEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Sprintf("Course with id %v not found", courseId)
		return []domain.WebinarSession{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return []domain.WebinarSession{}, err
	}

	webinarSessionEntities := courseEntity.Webinars
	err = db.WithContext(ctx).
		Where("course_id = ?", courseId).
		Preload("Day").
		Preload("Sequence").
		Find(&webinarSessionEntities).Error
	if err != nil {
		return []domain.WebinarSession{}, nil
	}

	var webinarSessions []domain.WebinarSession
	for _, webinarSession := range webinarSessionEntities {
		webinarSessions = append(webinarSessions, domain.WebinarSession{
			Id:             webinarSession.ID,
			CourseId:       webinarSession.CourseId,
			Title:          webinarSession.Title,
			Desc:           webinarSession.Desc.String,
			Cover:          webinarSession.Cover,
			ZoomLink:       webinarSession.ZoomLink,
			ScheduleDay:    webinarSession.Day.Name,
			ScheduleDate:   webinarSession.ScheduleDate,
			TimeStart:      webinarSession.TimeStart,
			TimeFinish:     webinarSession.TimeFinish,
			IsPublished:    webinarSession.IsPublished,
			SequenceNumber: webinarSession.Sequence.OrderNumber,
		})
	}

	return webinarSessions, nil
}

func (repository *WebinarSessionRepositoryImpl) GetDetailWebinarSessionByWebinarSessionId(ctx context.Context, db *gorm.DB, courseId uint, webinarSessionId uint) (domain.WebinarSession, error) {
	fmt.Println("Repository Get Detail Webinar Session")

	var courseEntity entity.MasterCourse
	err := db.WithContext(ctx).
		Where("id = ?", courseId).
		Preload("Webinars").
		First(&courseEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Sprintf("Course with id %v not found", courseId)
		return domain.WebinarSession{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return domain.WebinarSession{}, err
	}

	webinarSessionEntities := courseEntity.Webinars
	err = db.WithContext(ctx).
		Where("course_id = ?", courseId).
		Where("id = ?", webinarSessionId).
		Preload("Day").
		Preload("Sequence").
		First(&webinarSessionEntities).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) {
		logError := fmt.Sprintf("Webinar session with id %v not found", webinarSessionId)
		return domain.WebinarSession{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return domain.WebinarSession{}, err
	}

	webinarSessionEntity := webinarSessionEntities[0]

	return domain.WebinarSession{
		Id:             webinarSessionEntity.ID,
		CourseId:       webinarSessionEntity.CourseId,
		Title:          webinarSessionEntity.Title,
		Desc:           webinarSessionEntity.Desc.String,
		Cover:          webinarSessionEntity.Cover,
		ZoomLink:       webinarSessionEntity.ZoomLink,
		ScheduleDay:    webinarSessionEntity.Day.Name,
		ScheduleDate:   webinarSessionEntity.ScheduleDate,
		TimeStart:      webinarSessionEntity.TimeStart,
		TimeFinish:     webinarSessionEntity.TimeFinish,
		IsPublished:    webinarSessionEntity.IsPublished,
		SequenceNumber: webinarSessionEntity.Sequence.OrderNumber,
	}, nil
}
