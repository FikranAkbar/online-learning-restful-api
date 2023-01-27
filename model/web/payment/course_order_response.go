package payment

type CourseOrderResponse struct {
	OrderId     string               `json:"order_id" mapstructure:"user_id"`
	ItemDetails []ItemDetailResponse `json:"item_details" mapstructure:"item_details"`
}
