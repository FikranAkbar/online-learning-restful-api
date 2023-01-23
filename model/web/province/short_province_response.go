package province

type ShortProvinceResponse struct {
	Id           uint   `json:"id" mapstructure:"id"`
	ProvinceName string `json:"province_name" mapstructure:"province_name"`
}
