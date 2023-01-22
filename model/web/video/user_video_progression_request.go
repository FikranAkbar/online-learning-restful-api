package video

type UserVideoProgressionRequest struct {
	Progression float32 `json:"progression,omitempty"`
	IsComplete  bool    `json:"is_complete,omitempty"`
}
