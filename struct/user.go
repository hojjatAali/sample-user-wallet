package structs

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCreateRQ struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UserUpdateRQ struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
