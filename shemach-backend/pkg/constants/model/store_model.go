package model

// Store
type Store struct {
	ID              uint64   `json:"id"`
	OwnerID         uint64   `json:"owner_id"`
	AddressRef      uint64   `json:"-"`
	Address         *Address `json:"address"`
	ActiveProducts  uint     `json:"active_products"`
	StoreName       string   `json:"store_name"`
	ActiveContracts uint     `json:"active_contracts"`
	CreatedAt       uint64   `json:"created_at"`
	CreatedBy       uint64   `json:"created_by"`
}

// func (store *Store)  Format() {
// 	if store.Zone.
// }
