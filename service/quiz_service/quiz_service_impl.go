package quiz_service

import (
	"context"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/web/quiz"
	"online-learning-restful-api/repository/quiz_repository"
)

type QuizServiceImpl struct {
	quiz_repository.QuizRepository
	*gorm.DB
}

func NewQuizServiceImpl(repository quiz_repository.QuizRepository, DB *gorm.DB) *QuizServiceImpl {
	return &QuizServiceImpl{QuizRepository: repository, DB: DB}
}

func (service QuizServiceImpl) GetQuizAnswersByModuleId(ctx context.Context, courseId uint, moduleId uint) []quiz.DetailQuizAnswerResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	quizAnswers, err := service.QuizRepository.GetQuizAnswersByModuleId(ctx, tx, courseId, moduleId)
	helper.PanicIfError(err)

	var quizAnswersResponse []quiz.DetailQuizAnswerResponse
	for _, quizAnswer := range quizAnswers {
		quizAnswersResponse = append(quizAnswersResponse, quiz.DetailQuizAnswerResponse{
			QuizId:     quizAnswer.QuizId,
			UserId:     quizAnswer.UserId,
			QuizAnswer: quizAnswer.QuizAnswer,
			TotalLike:  quizAnswer.TotalLike,
			IsLiked:    quizAnswer.IsLiked,
			UserName:   quizAnswer.UserName,
			UserPhoto:  quizAnswer.UserPhoto,
		})
	}

	return quizAnswersResponse
}

