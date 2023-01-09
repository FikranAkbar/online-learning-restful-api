package entity

import (
	"gorm.io/gorm"
)

type TrxCourseQnaAnswer struct {
	gorm.Model    `gorm:"embedded"`
	QnaQuestionId uint                 `gorm:"column:qna_question_id;not null"`
	QnaQuestion   TrxCourseQnaQuestion `gorm:"foreignKey:QnaQuestionId;joinForeignKey:ID"`
	UserId        uint                 `gorm:"column:user_id;not null"`
	User          MasterUser           `gorm:"foreignKey:UserId;joinForeignKey:ID"`
	UserTypeId    uint                 `gorm:"column:user_type_id;not null"`
	UserType      MasterUserType       `gorm:"foreignKey:UserTypeId;joinForeignKey:ID"`
	Answer        string               `gorm:"column:answer;not null"`
}

func (TrxCourseQnaAnswer) TableName() string {
	return "trx_course_qna_answer"
}
