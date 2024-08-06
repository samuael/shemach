package model

// Img ...
type PostImg struct {
	ID             int    `json:"img"`
	Resource       string `json:"resource"`
	OwnerID        int    `json:"owner_id"`
	OwnerRole      int    `json:"owner_role"`
	Authorized     bool   `json:"authorized"`
	Authorizations int    `json:"authorizations"`
	CreatedAt      uint64 `json:"created_at"`
	BlurredRe      string `json:"blurred_res"`
}
