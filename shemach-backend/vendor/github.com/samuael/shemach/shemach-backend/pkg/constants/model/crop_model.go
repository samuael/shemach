package model

// Crop create a crop instance
type Crop struct {
	ID                uint64   `json:"id"`
	TypeID            uint     `json:"type_id"`
	RemainingQuantity uint64   `json:"remaining_quantity"`
	Description       string   `json:"description"`
	Negotiable        bool     `json:"negotiable"`
	SellingPrice      float64  `json:"selling_price"`
	Address           *Address `json:"address"`
	AddressRef        uint64   `json:"-"`
	Images            []int    `json:"images"`
	CreatedAt         uint64   `json:"created_at"`
	StoreID           uint64   `json:"store_id"`
	AgentID           uint64   `json:"agent_id"`
	StoreOwned        bool     `json:"store_owned"`
	Closed            bool     `json:"closed"`
}
