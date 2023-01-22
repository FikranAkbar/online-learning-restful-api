package quiz_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"online-learning-restful-api/exception"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"online-learning-restful-api/model/web/quiz"
	"online-learning-restful-api/repository/quiz_repository"
)

type QuizServiceImpl struct {
	quiz_repository.QuizRepository
	*validator.Validate
	*gorm.DB
}

func NewQuizServiceImpl(repository quiz_repository.QuizRepository, validate *validator.Validate, DB *gorm.DB) *QuizServiceImpl {
	return &QuizServiceImpl{QuizRepository: repository, Validate: validate, DB: DB}
}

func (service *QuizServiceImpl) GetQuizAnswersByModuleId(ctx context.Context, courseId uint, moduleId uint) []quiz.DetailQuizAnswerResponse {
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

func (service *QuizServiceImpl) CreateNewQuizAnswer(ctx context.Context, courseId uint, moduleId uint, request quiz.ShortQuizAnswerRequest) quiz.ShortQuizAnswerResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(exception.GenerateHTTPError(exception.BadRequest, err.Error()))
	}

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	quizAnswer := domain.QuizAnswer{QuizAnswer: request.Answer}

	quizAnswer, err = service.QuizRepository.CreateNewQuizAnswer(ctx, tx, courseId, moduleId, quizAnswer)
	helper.PanicIfError(err)

	return quiz.ShortQuizAnswerResponse{
		QuizId:     quizAnswer.QuizId,
		UserId:     quizAnswer.UserId,
		QuizAnswer: quizAnswer.QuizAnswer,
	}
}
