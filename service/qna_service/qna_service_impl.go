package qna_service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"online-learning-restful-api/helper"
	"online-learning-restful-api/model/domain"
	"online-learning-restful-api/model/web/qna"
	"online-learning-restful-api/repository/qna_repository"
)

type QnaServiceImpl struct {
	qna_repository.QnaRepository
	*validator.Validate
	*gorm.DB
}

func NewQnaServiceImpl(qnaRepository qna_repository.QnaRepository, validate *validator.Validate, DB *gorm.DB) *QnaServiceImpl {
	return &QnaServiceImpl{QnaRepository: qnaRepository, Validate: validate, DB: DB}
}

func (service *QnaServiceImpl) GetQnaQuestionsByCourseId(ctx context.Context, courseId uint) []qna.DetailQnaQuestionResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	qnaQuestions, users, err := service.QnaRepository.GetQnaQuestionsByCourseId(ctx, tx, courseId)
	helper.PanicIfError(err)

	var qnaQuestionsResponse []qna.DetailQnaQuestionResponse
	for i, qnaQuestion := range qnaQuestions {
		qnaQuestionsResponse = append(qnaQuestionsResponse, qna.DetailQnaQuestionResponse{
			QuestionId:  qnaQuestion.Id,
			Question:    qnaQuestion.Question,
			Responses:   qnaQuestion.Responses,
			CreatedTime: qnaQuestion.CreatedAt,
			UserName:    users[i].Name,
			UserPhoto:   users[i].PhotoURL,
		})
	}

	return qnaQuestionsResponse
}

func (service *QnaServiceImpl) CreateNewQnaQuestion(ctx context.Context, courseId uint, request qna.CreateQnaQuestionRequest) qna.ShortQnaQuestionResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	qnaQuestion := domain.QnaQuestion{
		Question: request.Question,
	}

	qnaQuestion, err := service.QnaRepository.CreateNewQnaQuestion(ctx, tx, courseId, qnaQuestion)
	helper.PanicIfError(err)

	return qna.ShortQnaQuestionResponse{
		QuestionId: qnaQuestion.Id,
		Question:   qnaQuestion.Question,
	}
}
