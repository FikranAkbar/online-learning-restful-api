package qna

import "time"

type DetailQnaAnswerResponse struct {
	AnswerId    uint      `json:"answer_id,omitempty" mapstructure:"answer_id"`
	Answer      string    `json:"answer,omitempty" mapstructure:"answer"`
	CreatedTime time.Time `json:"created_time" mapstructure:"created_time"`
	UserId      uint      `json:"user_id,omitempty" mapstructure:"user_id"`
	UserName    string    `json:"user_name,omitempty" mapstructure:"user_name"`
	UserPhoto   string    `json:"user_photo,omitempty" mapstructure:"user_photo"`
}
