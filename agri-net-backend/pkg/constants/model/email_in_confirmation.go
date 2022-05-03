package model

import (
	"encoding/json"
	"time"
)

type EmailInConfirmation struct {
	Email      string    `json:"email"`
	OldEmail   string    `json:"oldEmail"`
	CreatedAt  time.Time `json:"created_at"`
	NewAccount bool      `json:"newAccount"`
}

func (a *EmailInConfirmation) UnmarshalJSON(data []byte) error {
	type Alias EmailInConfirmation
	chil := &struct {
		*Alias
		CreatedAt int64 `json:"created_at"`
	}{
		Alias: (*Alias)(a),
	}
	if er := json.Unmarshal(data, chil); er != nil {
		return er
	}
	if chil.CreatedAt > 0 {
		a.CreatedAt = time.Unix(chil.CreatedAt, 0)
	}
	return nil
}
