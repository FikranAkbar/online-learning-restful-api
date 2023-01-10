package entity

import (
	"gorm.io/gorm"
)

type MasterInsightAttachment struct {
	gorm.Model     `gorm:"embedded"`
	AttachmentName string `gorm:"column:attachment_name;type:varchar(100);not null"`
	DocURL         string `gorm:"column:doc_url;default:http://www.africau.edu/images/default/sample.pdf"`
}

func (MasterInsightAttachment) TableName() string {
	return "master_insight_attachment"
}
