package payment

type ItemDetailResponse struct {
	Id    uint   `json:"id" mapstructure:"id"`
	Name  string `json:"name" mapstructure:"name"`
	Price uint   `json:"price" mapstructure:"price"`
}
