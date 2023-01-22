package industry_insight

type InsightAttachmentResponse struct {
	Id             uint   `json:"id" mapstructure:"id"`
	AttachmentName string `json:"attachment_name" mapstructure:"attachment_name"`
	DocURL         string `json:"doc_url" mapstructure:"doc_url"`
}
