package qna

type ShortQnaQuestionResponse struct {
	QuestionId uint   `json:"question_id" mapstructure:"question_id"`
	Question   string `json:"question" mapstructure:"question"`
}
