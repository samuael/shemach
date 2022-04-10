package model

type Superadmin struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
