package model

// Garage ...
type Garage struct {
	ID      uint     `csv:"id"  json:"id"`
	Name    string   `csv:"name"  json:"name"`
	Address *Address `csv:"address"  json:"address"`
}

// Address ...
type Address struct {
	ID      uint   `json:"id"`
	Country string `json:"country"`
	Region  string `json:"region"`
	Zone    string `json:"zone"`
	Woreda  string `json:"woreda"`
	City    string `json:"city"`
	Kebele  string `json:"kebele"`
}
