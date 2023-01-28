package payment

type MidtransPaymentNotificationRequest struct {
	TransactionStatus string `json:"transaction_status" mapstructure:"transaction_status"`
	OrderId           string `json:"order_id" mapstructure:"order_id"`
	PaymentType       string `json:"payment_type" mapstructure:"payment_type"`
}
