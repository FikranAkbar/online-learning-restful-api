package entity

import "online-learning-restful-api/core"

type TrxCourseQnaAnswer struct {
	core.EntityModel `gorm:"embedded"`
	QnaQuestionId    uint           `gorm:"column:qna_question_id;not null"`
	UserId           uint           `gorm:"column:user_id;not null"`
	MasterUser       MasterUser     `gorm:"foreignKey:UserId"`
	UserType         uint           `gorm:"column:user_type;not null"`
	MasterUserType   MasterUserType `gorm:"foreignKey:UserType"`
	Answer           string         `gorm:"column:answer;not null"`
}

func (TrxCourseQnaAnswer) TableName() string {
	return "trx_course_qna_answer"
}
