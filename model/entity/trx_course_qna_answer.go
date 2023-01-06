package entity

import "online-learning-restful-api/core"

type TrxCourseQnaAnswer struct {
	core.EntityModel
	QnaQuestionId uint `gorm:"not null"`
	UserId        uint `gorm:"not null"`
	UserType      uint `gorm:"not null"`
	Answer        string
}

func (TrxCourseQnaAnswer) TableName() string {
	return "trx_course_qna_answer"
}
