package qna

import "time"

type DetailQnaQuestionResponse struct {
	QuestionId  uint                      `json:"question_id" mapstructure:"question_id"`
	Question    string                    `json:"question" mapstructure:"question"`
	CreatedTime time.Time                 `json:"created_time" mapstructure:"created_time"`
	UserName    string                    `json:"user_name" mapstructure:"user_name"`
	UserPhoto   string                    `json:"user_photo" mapstructure:"user_photo"`
	Responses   int                       `json:"responses" mapstructure:"responses"`
	Answers     []DetailQnaAnswerResponse `json:"answers" mapstructure:"answers"`
}
