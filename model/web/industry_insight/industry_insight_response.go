package industry_insight

type IndustryInsightResponse struct {
	Id                 uint                        `json:"id,omitempty" mapstructure:"id"`
	TitleInsight       string                      `json:"title_insight,omitempty" mapstructure:"title_insight"`
	CoverInsight       string                      `json:"cover_insight,omitempty" mapstructure:"cover_insight"`
	BodyContent        string                      `json:"body_content,omitempty" mapstructure:"body_content"`
	IdUserAuthor       uint                        `json:"id_user_author,omitempty" mapstructure:"id_user_author"`
	InsightAttachments []InsightAttachmentResponse `json:"insight_attachments,omitempty" mapstructure:"insight_attachments"`
}
