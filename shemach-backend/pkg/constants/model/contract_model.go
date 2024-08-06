package model

type Contract struct {
	ID            uint64 `json:"id"`
	TransactionID uint64 `json:"transaction_id"`
	SecretString  string `json:"secret"`
	State         uint8  `json:"state"`
}
