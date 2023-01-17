package authentication

type UserRegisterResponse struct {
	Id    uint   `json:"id,omitempty" mapstructure:"id"`
	Name  string `json:"name,omitempty" mapstructure:"name"`
	Email string `json:"email,omitempty" mapstructure:"email"`
}
