package model

type EmailConfirmation struct {
	ID           uint64 `json:"id"`
	UserID       uint64 `json:"user_id"`
	Email        string `json:"email"`
	OldEmail     string `json:"old_email"`
	CreatedAt    uint64 `json:"created_at"`
	IsNewAccount bool   `json:"is_new_account"`
}
