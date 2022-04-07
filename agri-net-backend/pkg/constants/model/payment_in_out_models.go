package model

type PayinInput struct {
	Amount     float64 `json:"amount"`
	PayedBy    int     `json:"student_id"`
	UnixTime   int64   `json:"unix_timestamp"`
	CreatedAt  *Date   `json:"-"`
	RoundID    int     `json:"round_id"`
	UChars     string  `json:"uchars"`
	ReceivedBy int     `json:"-"`
}
