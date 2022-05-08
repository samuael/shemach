package model

// TempoCXP ...
type TempoCXP struct {
	ID           int    `json:"id"`
	Phone        string `json:"phone"`
	Confirmation string `json:"confirmation"`
	CreatedAt    uint64 `json:"created_at"`
	Role         int    `json:"role"`
}
