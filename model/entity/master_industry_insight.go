package entity

import "online-learning-restful-api/core"

type MasterIndustryInsight struct {
	core.EntityModel
	UserAuthorId uint   `gorm:"not null"`
	TitleInsight string `gorm:"type:varchar(100)"`
	CoverInsight string
	BodyContent  string
	IsPublished  bool `gorm:"default:false"`
}

func (MasterIndustryInsight) TableName() string {
	return "master_industry_insight"
}
