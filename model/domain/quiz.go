package domain

import "time"

type QuizQuestion struct {
	Id           uint
	ModuleId     uint
	QuizQuestion string
	IsPublished  bool
}

type QuizAnswer struct {
	QuizId     uint
	UserId     uint
	UserName   string
	UserPhoto  string
	QuizAnswer string
	IsLiked    bool
	TotalLike  uint
	CreatedAt  time.Time
}
