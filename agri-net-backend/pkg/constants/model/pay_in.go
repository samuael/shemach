package model

type PayIn struct {
	ID                uint    `json:"id"`
	Amount            float64 `json:"amount"`
	RecievedBy        uint    `json:"recieved_by"`
	StudentID         uint    `json:"student_id"`
	CreatedAt         *Date   `json:"created_at"`
	CreatedAtRef      uint    `json:"-"`
	RoundID           uint    `json:"round_id"`
	Status            uint8   `json:"status"`
	StatusDescription string  `json:"status_description,omitempty"`
	UChars            string  `json:"uchars"`
}
