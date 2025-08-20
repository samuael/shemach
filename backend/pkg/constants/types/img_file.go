package types

// Img ...
type PostImg struct {
	ID        int    `json:"img"`
	Role      int    `json:"role"`
	Resource  string `json:"resource"`
	CreatedBy uint64 `json:"created_by"`
	CreatedAt uint64 `json:"created_at"`

	// Optional blurred image id.
	BlurredPath string `json:"blurred_path"`
}
