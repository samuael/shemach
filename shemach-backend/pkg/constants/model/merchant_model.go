package model

// Merchant
type Merchant struct {
	User
	StoresCount   uint     `json:"stores_count"`
	PostsCount    uint     `json:"posts_count"`
	RegisteredBy  uint64   `json:"registered_by"`
	AddressRef    int64    `json:"-"`
	Address       *Address `json:"address"`
	Subscriptions []int    `json:"subscriptions"`
}
