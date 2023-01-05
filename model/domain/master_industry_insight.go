package domain

import "online-learning-restful-api/core"

type MasterIndustryInsight struct {
	core.DomainModel
	UserAuthorId uint   `gorm:"not null"`
	TitleInsight string `gorm:"type:varchar(100)"`
	CoverInsight string
	BodyContent  string
	IsPublished  bool `gorm:"default:false"`
}
