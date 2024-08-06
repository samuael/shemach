package model

type Dictionary struct {
	ID          uint64 `json:"id,omitempty"`
	Lang        string `json:"lang"`
	Text        string `json:"text"`
	Translation string `json:"translation"`
}
