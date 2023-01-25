package video

type UserVideoProgressionRequest struct {
	Progression float32 `json:"progression" validate:"required"`
	IsComplete  bool    `json:"is_complete"`
}
