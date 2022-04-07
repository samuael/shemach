package model

// Round represents the
type Round struct {
	ID           uint    `json:"id"` // id
	CategoryID   uint    `json:"category_id"`
	TrainingHour uint    `json:"training_hour"`
	RoundNo      uint    `json:"round_no"` // category
	Students     uint    `json:"students"` // students
	ActiveAmount float64 `json:"active_amount"`
	Active       bool    `json:"active"` // active
	StartDate    string  `json:"start_date"`
	Lang         string  `json:"lang"` // active_amount
	EndDate      string  `json:"end_date"`
	Fee          float64 `json:"fee"`
	CreatedAt    *Date   `json:"created_at"`
	CreatedAtRef uint    `json:"created_at_ref"`
}

/*
	If the Start Date and the End Date is not specified
	the handler must be able to respond with a status code
	or return an error message.
*/

// Changing the category of a single round is not allowed
type RoundInput struct {
	ID uint `json:"id"` // id
	// CategoryID   uint   `json:"category_id,omitempty"`
	TrainingHour uint `json:"training_hour,omitempty"`
	RoundNo      uint `json:"round_no,omitempty"` // category
	// Students     uint   `json:"students,omitempty"` // students
	StartDate string  `json:"start_date,omitempty"`
	EndDate   string  `json:"end_date,omitempty"`
	Lang      string  `json:"lang,omitempty"` // active_amount
	Fee       float64 `json:"fee,omitempty"`  // fee
}
