package model

// BirthDate in ethiopian calander only
type Date struct {
	ID       uint   `json:"id,omitempty"`
	Years    uint   `json:"year"`
	Months   uint8  `json:"month"`
	Days     uint8  `json:"day"`
	Hours    uint8  `json:"hours,omitempty"`
	Minutes  uint8  `json:"minutes,omitempty"`
	Seconds  uint8  `json:"seconds,omitempty"`
	DayName  string `json:"day_name,omitempty"`
	YearName string `json:"year_name,omitempty"`
}
