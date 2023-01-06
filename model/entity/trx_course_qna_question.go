package entity

import "online-learning-restful-api/core"

type TrxCourseQnaQuestion struct {
	core.EntityModel `gorm:"embedded"`
	CourseId         uint   `gorm:"not null"`
	UserId           uint   `gorm:"not null"`
	Question         string `gorm:"not null"`
}

func (TrxCourseQnaQuestion) TableName() string {
	return "trx_course_qna_question"
}
