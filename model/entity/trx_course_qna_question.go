package entity

import (
	"gorm.io/gorm"
)

type TrxCourseQnaQuestion struct {
	gorm.Model `gorm:"embedded"`
	CourseId   uint   `gorm:"column:course_id;not null"`
	UserId     uint   `gorm:"column:user_id;not null"`
	Question   string `gorm:"column:question;not null"`
}

func (TrxCourseQnaQuestion) TableName() string {
	return "trx_course_qna_question"
}
