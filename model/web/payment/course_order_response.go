package payment

type CourseOrderResponse struct {
	OrderId     uint                 `json:"user_id" mapstructure:"user_id"`
	ItemDetails []ItemDetailResponse `json:"item_details" mapstructure:"item_details"`
}
