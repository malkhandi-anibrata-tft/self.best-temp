package user

type UserDetailsinfo struct {
	UserId    uint64 `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateDetailsRequest struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
type UpdateDetailsRequest struct {
	UserId    uint64 `param:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type DeleteUserRequest struct {
	Email string `json:"email" validate:"required"`
}
