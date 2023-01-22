package domain

type ElearningModule struct {
	ID                uint
	CourseId          uint
	ModuleTitle       string
	ModuleOverview    string
	ModuleDescription string
	IsPublished       bool
	IsQuizAnswered    bool
	SequenceNumber    uint
}
