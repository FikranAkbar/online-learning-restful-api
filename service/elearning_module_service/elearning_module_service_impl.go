package elearning_module_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"online-learning-restful-api/model/web/elearning_module"
	"online-learning-restful-api/model/web/quiz"
	"online-learning-restful-api/model/web/video"
	"online-learning-restful-api/repository/elearning_module_repository"
	"sort"
)

type ElearningModuleServiceImpl struct {
	elearning_module_repository.ElearningModuleRepository
	*validator.Validate
	*gorm.DB
}

func NewElearningModuleServiceImpl(repository elearning_module_repository.ElearningModuleRepository, validate *validator.Validate, DB *gorm.DB) *ElearningModuleServiceImpl {
	return &ElearningModuleServiceImpl{ElearningModuleRepository: repository, Validate: validate, DB: DB}
}

func (service *ElearningModuleServiceImpl) GetOverviewElearningModulesByCourseId(ctx context.Context, courseId uint) []elearning_module.OverviewElearningModuleResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	elearningModules, err := service.ElearningModuleRepository.GetOverviewElearningModulesByCourseId(ctx, tx, courseId)
	helper.PanicIfError(err)

	var overviewElearningModulesResponse []elearning_module.OverviewElearningModuleResponse
	for _, elearningModule := range elearningModules {
		overviewElearningModulesResponse = append(overviewElearningModulesResponse, elearning_module.OverviewElearningModuleResponse{
			CourseId: elearningModule.CourseId,
			ModuleId: elearningModule.ID,
			Title:    elearningModule.ModuleTitle,
			Overview: elearningModule.ModuleOverview,
		})
	}

	return overviewElearningModulesResponse
}

func (service *ElearningModuleServiceImpl) GetDetailElearningModuleByElearningModuleId(ctx context.Context, courseId uint, elearningModuleId uint) elearning_module.DetailElearningModuleResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	elearningModule, videoData, quizQuestion, quizAnswers, err := service.ElearningModuleRepository.GetDetailElearningModuleByElearningModuleId(ctx, tx, courseId, elearningModuleId)
	helper.PanicIfError(err)

	sort.Slice(quizAnswers, func(i, j int) bool {
		return quizAnswers[i].CreatedAt.After(quizAnswers[j].CreatedAt)
	})
	var quizAnswersResponse []quiz.DetailQuizAnswerResponse
	for _, quizAnswer := range quizAnswers {
		quizAnswersResponse = append(quizAnswersResponse, quiz.DetailQuizAnswerResponse{
			UserId:     quizAnswer.UserId,
			QuizId:     quizAnswer.QuizId,
			QuizAnswer: quizAnswer.QuizAnswer,
			TotalLike:  quizAnswer.TotalLike,
			IsLiked:    quizAnswer.IsLiked,
			UserName:   quizAnswer.UserName,
			UserPhoto:  quizAnswer.UserPhoto,
		})
	}

	return elearning_module.DetailElearningModuleResponse{
		ModuleId:    elearningModule.ID,
		Title:       elearningModule.ModuleTitle,
		Overview:    elearningModule.ModuleOverview,
		Description: elearningModule.ModuleDescription,
		Video: video.DetailVideoResponse{
			VideoId:              videoData.Id,
			Name:                 videoData.VideoName,
			Url:                  videoData.VideoUrl,
			LastVideoProgression: videoData.LastVideoProgression,
			CourseId:             videoData.CourseId,
			ModuleId:             videoData.ModuleId,
		},
		IsAnswered: elearningModule.IsQuizAnswered,
		Quiz: quiz.DetailQuizQuestionResponse{
			QuizId:   quizQuestion.Id,
			Question: quizQuestion.QuizQuestion,
			Answers:  quizAnswersResponse,
		},
	}
}

func (service *ElearningModuleServiceImpl) SaveVideoProgressionInModule(ctx context.Context, courseId uint, elearningModuleId uint, request video.UserVideoProgressionRequest) video.UserVideoProgressionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	data := domain.UserVideoProgression{
		Progression: request.Progression,
		IsComplete:  request.IsComplete,
	}

	userVideoProgression, err := service.ElearningModuleRepository.SaveVideoProgressionInModule(ctx, tx, courseId, elearningModuleId, data)
	helper.PanicIfError(err)

	userVideoProgressionResponse := video.UserVideoProgressionResponse{
		VideoId:      userVideoProgression.VideoId,
		UserId:       userVideoProgression.UserId,
		Progressions: userVideoProgression.Progression,
		IsComplete:   userVideoProgression.IsComplete,
	}

	return userVideoProgressionResponse
}
