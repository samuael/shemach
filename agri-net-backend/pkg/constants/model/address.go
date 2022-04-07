package model

// Address ...
type Address struct {
	ID   uint   `json:"id,omitempty"`
	City string `json:"city,omitempty"`
	// length should be >=3
	Region string `json:"region"`
	// length should be >=3
	Zone string `json:"zone"`
	// length should be >=3
	Woreda            string `json:"woreda"`
	Kebele            string `json:"kebele"`
	UniqueAddressName string `json:"unique_address_name,omitempty"`
}
