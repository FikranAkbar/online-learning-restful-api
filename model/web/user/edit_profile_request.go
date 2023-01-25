package user

type EditProfileRequest struct {
	Name         string `json:"name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	BirthDate    string `json:"birth_date,omitempty"`
	Gender       string `json:"gender,omitempty"`
	Photo        string `json:"photo,omitempty"`
	ProvinceName string `json:"province_name,omitempty"`
	CityName     string `json:"city_name,omitempty"`
}
