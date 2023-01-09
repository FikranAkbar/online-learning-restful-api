package entity

import (
	"gorm.io/gorm"
)

type TrxCourseQnaAnswer struct {
	gorm.Model    `gorm:"embedded"`
	QnaQuestionId uint   `gorm:"column:qna_question_id;not null"`
	UserId        uint   `gorm:"column:user_id;not null"`
	UserType      uint   `gorm:"column:user_type;not null"`
	Answer        string `gorm:"column:answer;not null"`
}

func (TrxCourseQnaAnswer) TableName() string {
	return "trx_course_qna_answer"
}
