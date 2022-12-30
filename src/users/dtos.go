package users

type LoginDTO struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
