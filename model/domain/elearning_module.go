package domain

type ElearningModule struct {
	ID             uint
	CourseId       uint
	ModuleTitle    string
	ModuleOverview string
	IsPublished    bool
	SequenceNumber uint
}
