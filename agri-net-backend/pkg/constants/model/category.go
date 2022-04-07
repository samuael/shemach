package model

// Category represents the categories.
type Category struct {
	ID           uint64  `json:"id"`
	Imageurl     string  `json:"imgurl"`
	Title        string  `json:"title"`
	ShortTitle   string  `json:"short_title"`
	RoundsCount  int     `json:"rounds_count"`
	Fee          float64 `json:"fee"`
	CreatedAtRef uint    `json:"-"`
	CreatedAt    *Date   `json:"created_at"`
}
