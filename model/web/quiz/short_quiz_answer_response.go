package quiz

type ShortQuizAnswerResponse struct {
	QuizId     uint   `json:"quiz_id" mapstructure:"quiz_id"`
	UserId     uint   `json:"user_id" mapstructure:"user_id"`
	QuizAnswer string `json:"quiz_answer" mapstructure:"quiz_answer"`
}
