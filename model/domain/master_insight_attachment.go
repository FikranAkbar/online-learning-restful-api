package domain

import "online-learning-restful-api/core"

type MasterInsightAttachment struct {
	core.DomainModel
	AttachmentName string `gorm:"type:varchar(100)"`
	DocURL         string
}
