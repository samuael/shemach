package model

type Superadmin struct {
	User
	CreatedAtUnix      int64 `json:"created_at_unix"`
	RegisteredAdmins   uint8 `json:"registered_admins"`
	RegisteredProducts uint8 `json:"registered_products"`
}
