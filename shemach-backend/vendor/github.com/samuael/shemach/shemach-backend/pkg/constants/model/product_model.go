package model

type Product struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	ProductionArea string  `json:"production_area"`
	UnitID         uint8   `json:"unit_id"`
	CurrentPrice   float64 `json:"current_price"`
	CreatedBy      uint64  `json:"created_by"`
	CreatedAt      uint64  `json:"created_at"`
	LastUpdateTime uint64  `json:"last_update_time"`
}

type ProductUpdate struct {
	ID    uint8   `json:"id"`
	Price float64 `json:"cost"`
}
