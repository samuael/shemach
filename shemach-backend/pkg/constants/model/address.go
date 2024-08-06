package model

// Address ...
type Address struct {
	ID                uint    `json:"id,omitempty"`
	Kebele            string  `json:"kebele"`
	Woreda            string  `json:"woreda"`
	City              string  `json:"city,omitempty"`
	UniqueAddressName string  `json:"unique_address,omitempty"`
	Region            string  `json:"region"`
	Zone              string  `json:"zone"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
}

/*

{
	"kebele":"04",
	"woreda":"01",
	"city":"Assosa",
	"unique_address":"Ambab Ber",
	"region":"Benishangul",
	"zone":"Assosa",
	"latitude":45898955.44343,
	"longitude":432423423423.99
}

*/
