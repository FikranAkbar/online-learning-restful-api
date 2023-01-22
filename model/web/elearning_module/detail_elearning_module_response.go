package elearning_module

import (
	"online-learning-restful-api/model/web/quiz"
	"online-learning-restful-api/model/web/video"
)

type DetailElearningModuleResponse struct {
	ModuleId    uint
	Title       string
	Overview    string
	Description string
	Video       video.DetailVideoResponse
	IsAnswered  bool
	Quiz        quiz.DetailQuizQuestionResponse
}
