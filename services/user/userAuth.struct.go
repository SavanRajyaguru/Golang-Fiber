package user

type (
	UserSignUp struct {
		Username string `json:"username" validate:"required,min=2"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}
)
