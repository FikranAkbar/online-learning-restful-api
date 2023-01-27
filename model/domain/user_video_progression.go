package domain

type UserVideoProgression struct {
	ID          uint
	VideoId     uint
	UserId      uint
	Progression float32
	IsComplete  bool
}
