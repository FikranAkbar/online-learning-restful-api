package entity

import "online-learning-restful-api/core"

type TrxCourseQnaQuestion struct {
	core.EntityModel `gorm:"embedded"`
	CourseId         uint               `gorm:"column:course_id;not null"`
	MasterCourse     MasterCourseStatus `gorm:"foreignKey:CourseId"`
	UserId           uint               `gorm:"column:user_id;not null"`
	MasterUser       MasterUser         `gorm:"foreignKey:UserId"`
	Question         string             `gorm:"column:question;not null"`
}

func (TrxCourseQnaQuestion) TableName() string {
	return "trx_course_qna_question"
}
