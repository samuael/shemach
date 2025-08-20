package types

// User model representing the admin, all of this instances has to satisfy this behaviour.
type User struct {
	ID             uint64 `json:"id,omitempty"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Phone          string `json:"phone,omitempty"`
	Email          string `json:"email"`
	Telegram       string `json:"telegramID,omitempty"`
	ProfileImageID uint64 `json:"profile_img_id"`
	CreatedAt      uint64 `json:"created_at,omitempty"`
	Bio            string `json:"bio"`
	Password       string `json:"-"`
	Role           uint8  `json:"role"`
}

// UserAuthData holds a tempo user auth data registration information
type UserAuthData struct {
	ID        uint64 `json:"id"`
	Phone     string `json:"phone"`
	CreatedAt int64  `json:"created_at"`
	Code      string `json:"code"`
	Trials    uint8  `json:"trials"`
}

// ConfirmationCode
type ConfirmationCode struct {
	ID    uint64 `json:"id"`
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// RegistrationData represents a registration details
type RegistrationData struct {
	Token      string `json:"token"`
	Email      string `json:"email"`
	TelegramID string `json:"telegram_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Password   string `json:"password"`
	Role       uint8  `json:"role"`
}
