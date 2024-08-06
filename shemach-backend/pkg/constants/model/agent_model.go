package model

type Agent struct {
	User
	PostsCount      uint     `json:"posts_count"`
	FieldAddressRef uint64   `json:"-"`
	FieldAddress    *Address `json:"field_address"`
	RegisteredBy    uint     `json:"registered_by"`
}
