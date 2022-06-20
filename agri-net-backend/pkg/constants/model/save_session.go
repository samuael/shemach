package model

type SaveSession struct {
	ID     int
	UserID int
	Token  string
}

type SaveSubscriberSession struct {
	ID           int
	SubscriberID int
	Token        string
}
