package model

type Product struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	ProductionArea string  `json:"production_area"`
	CurrentPrice   float64 `json:"current_price"`
	CreatedAt      uint64  `json:"created_at"`
	CreatedBy      uint64  `json:"created_by"`
	LastUpdateTime uint64  `json:"last_update_time"`
}
