package authentication

type UserLoginResponse struct {
	Email string `json:"email" mapstructure:"email"`
	Name  string `json:"name" mapstructure:"name"`
	Token string `json:"token" mapstructure:"token"`
}
