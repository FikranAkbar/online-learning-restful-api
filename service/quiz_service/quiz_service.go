package quiz_service

import (
	"context"
	"online-learning-restful-api/model/web/quiz"
)

type QuizService interface {
	GetQuizAnswersByModuleId(ctx context.Context, courseId uint, moduleId uint) []quiz.DetailQuizAnswerResponse
	CreateNewQuizAnswer(ctx context.Context, courseId uint, moduleId uint, request quiz.ShortQuizAnswerRequest) quiz.ShortQuizAnswerResponse
}
