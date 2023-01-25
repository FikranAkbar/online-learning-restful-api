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
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	qnaQuestion := domain.QnaQuestion{
		Question: request.Question,
	}

	qnaQuestion, err = service.QnaRepository.CreateNewQnaQuestion(ctx, tx, courseId, qnaQuestion)
	helper.PanicIfError(err)

	return qna.ShortQnaQuestionResponse{
		QuestionId: qnaQuestion.Id,
		Question:   qnaQuestion.Question,
	}
}

func (service *QnaServiceImpl) GetDetailQnaQuestionByQnaQuestionId(ctx context.Context, courseId uint, qnaQuestionId uint) qna.DetailQnaQuestionResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	qnaQuestion, qnaAnswers, err := service.QnaRepository.GetDetailQnaQuestionByQnaQuestionId(ctx, tx, courseId, qnaQuestionId)
	helper.PanicIfError(err)

	var qnaAnswersResponse []qna.DetailQnaAnswerResponse
	for _, qnaAnswer := range qnaAnswers {
		qnaAnswersResponse = append(qnaAnswersResponse, qna.DetailQnaAnswerResponse{
			AnswerId:    qnaAnswer.Id,
			Answer:      qnaAnswer.Answer,
			CreatedTime: qnaAnswer.CreatedAt,
			UserId:      qnaAnswer.UserId,
			UserName:    qnaAnswer.UserName,
			UserPhoto:   qnaAnswer.UserPhoto,
		})
	}

	return qna.DetailQnaQuestionResponse{
		QuestionId:  qnaQuestion.Id,
		Question:    qnaQuestion.Question,
		CreatedTime: qnaQuestion.CreatedAt,
		UserName:    qnaQuestion.UserName,
		UserPhoto:   qnaQuestion.UserPhoto,
		Responses:   len(qnaAnswersResponse),
		Answers:     qnaAnswersResponse,
	}
}

func (service *QnaServiceImpl) CreateNewQnaAnswer(ctx context.Context, courseId uint, qnaQuestionId uint, request qna.CreateQnaAnswerRequest) qna.ShortQnaAnswerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	qnaAnswer := domain.QnaAnswer{
		Answer: request.Answer,
	}

	qnaAnswer, err = service.QnaRepository.CreateNewQnaAnswer(ctx, tx, courseId, qnaQuestionId, qnaAnswer)
	helper.PanicIfError(err)

	return qna.ShortQnaAnswerResponse{
		AnswerId: qnaAnswer.Id,
		Answer:   qnaAnswer.Answer,
	}
}

