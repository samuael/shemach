package model

type Message struct {
	ID        uint64 `json:"id,omitempty"`
	Targets   []int  `json:"targets"`
	Lang      string `json:"lang"`
	Data      string `json:"data"`
	CreatedBy uint   `json:"created_by,omitempty"`
	CreatedAt uint64 `json:"created_at,omitempty"`
}
