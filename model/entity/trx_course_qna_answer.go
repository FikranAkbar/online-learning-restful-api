package entity

import "online-learning-restful-api/core"

type TrxCourseQnaAnswer struct {
	core.EntityModel `gorm:"embedded"`
	QnaQuestionId    uint           `gorm:"not null"`
	UserId           uint           `gorm:"not null"`
	MasterUser       MasterUser     `gorm:"UserId"`
	UserType         uint           `gorm:"not null"`
	MasterUserType   MasterUserType `gorm:"UserType"`
	Answer           string         `gorm:"not null"`
}

func (TrxCourseQnaAnswer) TableName() string {
	return "trx_course_qna_answer"
}
