package model

type ErMsg struct {
	Error  string `json:"err"`
	Status int    `json:"status,omitempty"`
}
