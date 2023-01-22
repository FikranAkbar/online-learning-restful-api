package quiz

type DetailQuizQuestionResponse struct {
	QuizId   uint
	Question string
	Answers  []DetailQuizAnswerResponse
}
