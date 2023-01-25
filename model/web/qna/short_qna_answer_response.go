package qna

type ShortQnaAnswerResponse struct {
	AnswerId uint   `json:"answer_id" mapstructure:"answer_id"`
	Answer   string `json:"answer" mapstructure:"answer"`
}
