package model

type PayOut struct {
	ID           uint    `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Amount       float64 `json:"amount"`
	Approved     bool    `json:"approved"`
	WithdrawedBy uint    `json:"withdrawed_by"`
	CreatedAt    *Date   `json:"created_at"`
	Status       uint8   `json:"status"`
	CreatedAtRef uint    `json:"-"`
}
