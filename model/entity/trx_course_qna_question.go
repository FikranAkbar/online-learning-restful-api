package entity

import "online-learning-restful-api/core"

type TrxCourseQnaQuestion struct {
	core.EntityModel
	CourseId uint `gorm:"not null"`
	UserId   uint `gorm:"not null"`
	Question string
}

func (TrxCourseQnaQuestion) TableName() string {
	return "trx_course_qna_question"
}
