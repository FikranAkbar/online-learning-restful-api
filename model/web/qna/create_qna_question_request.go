package qna

type CreateQnaQuestionRequest struct {
	Question string `json:"question,omitempty" validate:"required"`
}
