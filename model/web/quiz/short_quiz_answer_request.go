package quiz

type ShortQuizAnswerRequest struct {
	Answer string `json:"answer" validate:"required"`
}
