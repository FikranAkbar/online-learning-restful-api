package qna

type CreateQnaAnswerRequest struct {
	Answer string `json:"answer,omitempty" validate:"required"`
}
