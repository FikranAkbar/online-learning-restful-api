package entity

import (
	"time"
)

type MasterInsightAttachment struct {
	ID             uint `gorm:"column:id;primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	AttachmentName string `gorm:"column:attachment_name;type:varchar(100);not null"`
	DocURL         string `gorm:"column:doc_url;default:http://www.africau.edu/images/default/sample.pdf"`
}

func (MasterInsightAttachment) TableName() string {
	return "master_insight_attachment"
}
