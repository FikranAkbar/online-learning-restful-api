package authentication

type UserRegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=20"`
	Name      string `json:"name" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
}
