package model

// User model representing the admin, info admins, and superadmins
// all of this instances has to satisfy this behaviour.
type User struct {
	ID        uint64 `json:"id,omitempty"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email"`
	Imgurl    string `json:"imgurl,omitempty"`
	CreatedAt uint64 `json:"created_at,omitempty"`
	Password  string `json:"password,omitempty"`
}

// InfoAdmin instance representing the product information supervisors.
type Infoadmin struct {
	User
	BroadcastedMessagesCount uint `json:"broadcasted_messages"`
	Createdby                uint64
}
