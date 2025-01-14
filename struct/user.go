package structs

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCreateRQ struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}
type UserUpdateRQ struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

func (u *UserCreateRQ) Validate() error {
	return validate.Struct(u)
}

func (u *UserUpdateRQ) Validate() error {
	return validate.Struct(u)
}
