package quiz

type DetailQuizAnswerResponse struct {
	QuizId     uint   `json:"quiz_id" mapstructure:"quiz_id"`
	UserId     uint   `json:"user_id" mapstructure:"user_id"`
	QuizAnswer string `json:"quiz_answer" mapstructure:"quiz_answer"`
	TotalLike  uint   `json:"total_like" mapstructure:"total_like"`
	IsLiked    bool   `json:"is_liked" mapstructure:"is_liked"`
	UserName   string `json:"user_name" mapstructure:"user_name"`
	UserPhoto  string `json:"user_photo" mapstructure:"user_photo"`
}
