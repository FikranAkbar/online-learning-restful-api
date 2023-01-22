package domain

type IndustryInsight struct {
	Id           uint
	UserAuthorId uint
	TitleInsight string
	CoverInsight string
	BodyContent  string
	IsPublished  bool
}
