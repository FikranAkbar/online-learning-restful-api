package quiz_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"time"
)

type QuizRepositoryImpl struct {
}

func NewQuizRepositoryImpl() *QuizRepositoryImpl {
	return &QuizRepositoryImpl{}
}

func (repository *QuizRepositoryImpl) GetQuizAnswersByModuleId(ctx context.Context, db *gorm.DB, courseId uint, elearningModuleId uint) ([]domain.QuizAnswer, error) {
	userTokenInfo, ok := ctx.Value(middleware.ContextUserInfoKey).(middleware.UserTokenInfo)
	if !ok {
		panic(middleware.UnauthorizedErrorInfo)
	}

	// Finding Entities Data
	var courseEntity entity.MasterCourse
	err := db.WithContext(ctx).
		Where("id = ?", courseId).
		Preload("Modules").
		First(&courseEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) || !courseEntity.IsPublished {
		logError := fmt.Sprintf("Course with id %v not found", courseId)
		return []domain.QuizAnswer{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return []domain.QuizAnswer{}, err
	}

	elearningModuleEntities := courseEntity.Modules
	err = db.WithContext(ctx).
		Where("course_id = ?", courseId).
		Where("id = ?", elearningModuleId).
		Preload("Quiz").
		Preload("Video").
		Preload("Sequence").
		First(&elearningModuleEntities).Error
	if (err != nil && exception.CheckErrorContains(err, exception.NotFound)) || !elearningModuleEntities[0].IsPublished {
		logError := fmt.Sprintf("E-learning Module with id %v not found", elearningModuleId)
		return []domain.QuizAnswer{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return []domain.QuizAnswer{}, err
	}

	elearningModuleEntity := elearningModuleEntities[0]

	videoEntity := elearningModuleEntity.Video
	err = db.WithContext(ctx).
		Where("course_id = ?", courseId).
		Where("module_id = ?", elearningModuleId).
		Preload("UserVideoProgression", "user_id = ?", userTokenInfo.UserId).
		Find(&videoEntity).Error
	helper.PanicIfError(err)

	quizEntity := elearningModuleEntity.Quiz
	err = db.WithContext(ctx).
		Where("module_id = ?", elearningModuleId).
		Preload("Answers").
		First(&quizEntity).Error
	if (err != nil && exception.CheckErrorContains(err, exception.NotFound)) || !quizEntity.IsPublished {
		logError := fmt.Sprintf("Quiz question in module with id %v not found", elearningModuleEntity.ID)
		return []domain.QuizAnswer{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return []domain.QuizAnswer{}, err
	}

	quizAnswerEntities := quizEntity.Answers
	err = db.WithContext(ctx).
		Where("quiz_id = ?", quizEntity.ID).
		Preload("User").
		Find(&quizAnswerEntities).Error
	helper.PanicIfError(err)
	//

	var quizAnswers []domain.QuizAnswer
	for _, quizAnswerEntity := range quizAnswerEntities {
		err = db.WithContext(ctx).
			Where("id = ?", quizAnswerEntity.ID).
			Where("quiz_id = ?", quizAnswerEntity.QuizId).
			Preload("User").
			First(&quizAnswerEntity).Error
		helper.PanicIfError(err)

		var isLikedByCurrentUser bool
		err = db.WithContext(ctx).
			Where("likeable_id = ?", quizAnswerEntity.ID).
			Where("likeable_type_id = ?", 1).
			Where("user_id = ?", userTokenInfo.UserId).
			Where("is_like = ?", true).
			First(&entity.TrxUserLike{}).Error
		if err == nil {
			isLikedByCurrentUser = true
		}

		var userLikes []entity.TrxUserLike
		err = db.WithContext(ctx).
			Where("likeable_id = ?", quizAnswerEntity.ID).
			Where("likeable_type_id = ?", 1).
			Where("is_like = ?", true).
			Find(&userLikes).Error
		helper.PanicIfError(err)

		quizAnswers = append(quizAnswers, domain.QuizAnswer{
			QuizId:     quizAnswerEntity.QuizId,
			UserId:     quizAnswerEntity.UserId,
			UserName:   quizAnswerEntity.User.Name,
			UserPhoto:  quizAnswerEntity.User.PhotoURL.String,
			QuizAnswer: quizAnswerEntity.QuizAnswer,
			IsLiked:    isLikedByCurrentUser,
			TotalLike:  uint(len(userLikes)),
			CreatedAt:  time.Time{},
		})
	}

	return quizAnswers, nil
}
