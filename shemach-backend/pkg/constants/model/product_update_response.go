package model

type ProductUpdateMessage struct {
	ID            uint    `json:"id"`
	PreviousPrice float64 `json:"previous_price"`
	NewPrice      float64 `json:"new_price"`
	When          uint64  `json:"when"`
}
