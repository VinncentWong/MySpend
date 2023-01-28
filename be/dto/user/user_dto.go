package user

type UserRegister struct {
	Name     string `validate:"required,min=6"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

type UserLogin struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
