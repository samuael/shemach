package model

// SpecialCase
type SpecialCase struct {
	StudentID uint64  `json:"-"`
	ID        uint    `json:"id"`
	Reason    string  `json:"reason"`
	Amount    float64 `json:"covered_amount,omitempty"`
	Total     bool    `json:"complete_fee"`
}
