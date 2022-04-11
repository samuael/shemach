package model

type Superadmin struct {
	ID                 uint64 `json:"id"`
	Firstname          string `json:"firstname"`
	Lastname           string `json:"lastname"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	Password           string `json:"password"`
	CreatedAtUnix      int64  `json:"created_at_unix"`
	CreatedAt          *Date  `json:"created_at,omitempty"`
	RegisteredAdmins   uint8  `json:"registered_admins"`
	RegisteredProducts uint8  `json:"registered_products"`
}
