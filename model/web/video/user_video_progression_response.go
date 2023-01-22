package video

type UserVideoProgressionResponse struct {
	VideoId      uint
	UserId       uint
	Progressions float32
	IsComplete   bool
}
