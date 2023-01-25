package qna_service

import (
	"context"
	"online-learning-restful-api/model/web/qna"
)

type QnaService interface {
	GetQnaQuestionsByCourseId(ctx context.Context, courseId uint) []qna.DetailQnaQuestionResponse
	CreateNewQnaQuestion(ctx context.Context, courseId uint, request qna.CreateQnaQuestionRequest) qna.ShortQnaQuestionResponse
	GetDetailQnaQuestionByQnaQuestionId(ctx context.Context, courseId uint, qnaQuestionId uint) qna.DetailQnaQuestionResponse
	CreateNewQnaAnswer(ctx context.Context, courseId uint, qnaQuestionId uint, request qna.CreateQnaAnswerRequest) qna.ShortQnaAnswerResponse
}
