package city

type ShortCityResponse struct {
	Id         uint   `json:"id" mapstructure:"id"`
	CityName   string `json:"city_name" mapstructure:"city_name"`
	ProvinceId uint   `json:"province_id" mapstructure:"province_id"`
}
