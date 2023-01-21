package elearning_module_repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
	"online-learning-restful-api/app/router/middleware"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"sort"
	"time"
)

type ElearningModuleRepositoryImpl struct {
}

func NewElearningModuleRepositoryImpl() *ElearningModuleRepositoryImpl {
	return &ElearningModuleRepositoryImpl{}
}

func (repository *ElearningModuleRepositoryImpl) GetOverviewElearningModulesByCourseId(ctx context.Context, db *gorm.DB, courseId uint) ([]domain.ElearningModule, error) {
	var courseEntity entity.MasterCourse
	err := db.WithContext(ctx).
		Where("id = ?", courseId).
		Preload("Modules").
		First(&courseEntity).Error
	if err != nil && exception.CheckErrorContains(err, exception.NotFound) || !courseEntity.IsPublished {
		logError := fmt.Sprintf("Course with id %v not found", courseId)
		return []domain.ElearningModule{}, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return []domain.ElearningModule{}, err
	}

	moduleEntities := courseEntity.Modules
	err = db.WithContext(ctx).
		Where("course_id = ?", courseId).
		Preload("Sequence").
		Find(&moduleEntities).Error
	if err != nil {
		return []domain.ElearningModule{}, nil
	}

	var modules []domain.ElearningModule
	for _, moduleEntity := range moduleEntities {
		if !moduleEntity.IsPublished {
			continue
		}

		fmt.Println("Order Number:", moduleEntity.Sequence.OrderNumber)
		modules = append(modules, domain.ElearningModule{
			ID:             moduleEntity.ID,
			CourseId:       moduleEntity.CourseId,
			ModuleTitle:    moduleEntity.ModuleTitle,
			ModuleOverview: moduleEntity.ModuleOverview.String,
			IsPublished:    moduleEntity.IsPublished,
			SequenceNumber: moduleEntity.Sequence.OrderNumber,
		})
	}

	sort.Slice(modules, func(i, j int) bool {
		return modules[i].SequenceNumber < modules[j].SequenceNumber
	})

	return modules, nil
}

func (repository *ElearningModuleRepositoryImpl) GetDetailElearningModuleByElearningModuleId(ctx context.Context, db *gorm.DB, courseId uint, elearningModuleId uint) (domain.ElearningModule, domain.Video, domain.QuizQuestion, []domain.QuizAnswer, error) {
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
		return domain.ElearningModule{}, domain.Video{}, domain.QuizQuestion{}, nil, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return domain.ElearningModule{}, domain.Video{}, domain.QuizQuestion{}, nil, err
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
		return domain.ElearningModule{}, domain.Video{}, domain.QuizQuestion{}, nil, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return domain.ElearningModule{}, domain.Video{}, domain.QuizQuestion{}, nil, err
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
		return domain.ElearningModule{}, domain.Video{}, domain.QuizQuestion{}, nil, exception.GenerateHTTPError(exception.NotFound, logError)
	} else if err != nil {
		return domain.ElearningModule{}, domain.Video{}, domain.QuizQuestion{}, nil, err
	}

	quizAnswerEntities := quizEntity.Answers
	err = db.WithContext(ctx).
		Where("quiz_id = ?", quizEntity.ID).
		Preload("User").
		Find(&quizAnswerEntities).Error
	helper.PanicIfError(err)
	//

	// Convert Entity to Domain
	var isQuizAnsweredByCurrentUser = false
	var quizAnswers []domain.QuizAnswer
	for _, quizAnswerEntity := range quizAnswerEntities {
		if quizAnswerEntity.UserId == userTokenInfo.UserId {
			isQuizAnsweredByCurrentUser = true
		}

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

	elearningModule := domain.ElearningModule{
		ID:                elearningModuleEntity.ID,
		CourseId:          elearningModuleEntity.CourseId,
		ModuleTitle:       elearningModuleEntity.ModuleTitle,
		ModuleOverview:    elearningModuleEntity.ModuleOverview.String,
		ModuleDescription: elearningModuleEntity.ModuleDescription.String,
		IsPublished:       elearningModuleEntity.IsPublished,
		IsQuizAnswered:    isQuizAnsweredByCurrentUser,
		SequenceNumber:    elearningModuleEntity.Sequence.OrderNumber,
	}

	video := domain.Video{
		Id:                   videoEntity.ID,
		CourseId:             videoEntity.CourseId,
		ModuleId:             videoEntity.ModuleId,
		VideoName:            videoEntity.VideoName,
		VideoUrl:             videoEntity.VideoUrl,
		LastVideoProgression: videoEntity.UserVideoProgression.Progression,
		IsPublished:          videoEntity.IsPublished,
	}

	quizQuestion := domain.QuizQuestion{
		Id:           quizEntity.ID,
		ModuleId:     quizEntity.ModuleId,
		QuizQuestion: quizEntity.QuizQuestion,
		IsPublished:  quizEntity.IsPublished,
	}
	//

	return elearningModule, video, quizQuestion, quizAnswers, nil
}
