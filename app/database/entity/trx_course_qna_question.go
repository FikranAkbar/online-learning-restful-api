package entity

import (
	"time"
)

type TrxCourseQnaQuestion struct {
	ID        uint `gorm:"column:id;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CourseId  uint                 `gorm:"column:course_id;not null"`
	UserId    uint                 `gorm:"column:user_id;not null"`
	User      MasterUser           `gorm:"foreignKey:UserId;joinForeignKey:ID"`
	Question  string               `gorm:"column:question;not null"`
	Answers   []TrxCourseQnaAnswer `gorm:"foreignKey:QnaQuestionId;joinForeignKey:ID"`
}

func (TrxCourseQnaQuestion) TableName() string {
	return "trx_course_qna_question"
}
