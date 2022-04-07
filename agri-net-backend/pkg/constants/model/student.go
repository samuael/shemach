package model

// Student instance.
type Student struct {
	ID              uint         `json:"id"`
	Fullname        string       `json:"fullname"`
	Sex             string       `json:"sex"`
	Age             uint         `json:"age,omitempty"`
	BirthDate       *Date        `json:"birth_date"`
	BirthDateRef    uint         `json:"-"`
	AccStatus       string       `json:"acc_status"`
	Address         *Address     `json:"address"`
	AddressRef      uint         `json:"-"`
	Phone           string       `json:"phone"`
	PaidAmount      float64      `json:"paid_amount,omitempty"`
	Status          uint8        `json:"status,omitempty"`
	RegisteredBy    uint         `json:"registered_by"`
	RoundID         uint         `json:"round_id"`
	Imgurl          string       `json:"imgurl,omitempty"`
	Marked          *SpecialCase `json:"special_case,omitempty"`
	MarkedRef       int          `json:"-"`
	RegisteredAtRef uint         `json:"-"`
	RegisteredAt    *Date        `json:"registered_at,omitempty"`
}
