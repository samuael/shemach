package model

// Admin represents all the admin and superadmin
type Admin struct {
	ID           uint   `json:"id"`
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	Imgurl       string `json:"imgurl"`
	Superadmin   bool   `json:"superadmin"`
	CreatedAtRef int    `json:"-"`
	CreatedAt    *Date  `json:"created_at"`
}
