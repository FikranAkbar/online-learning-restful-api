package user

import "time"

type DetailUserResponse struct {
	Id         uint      `json:"id,omitempty" mapstructure:"id"`
	Name       string    `json:"name,omitempty" mapstructure:"name"`
	BirthDate  time.Time `json:"birth_date" mapstructure:"birth_date"`
	Gender     string    `json:"gender,omitempty" mapstructure:"gender"`
	Phone      string    `json:"phone,omitempty" mapstructure:"phone"`
	PhotoURL   string    `json:"photo_url,omitempty" mapstructure:"photo_url"`
	ProvinceId uint      `json:"province_id,omitempty" mapstructure:"province_id"`
}
