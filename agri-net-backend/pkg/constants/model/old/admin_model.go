package model

// Admin a user representing  Shop Owner
type Admin struct {
	ID              uint64 `json:"id" 			gorm:"primaryKey;autoIncrement:true"`
	Email           string `json:"email"  		gormSigningKey:"type:varchar(255);not null; unique;"`
	Firstname       string `json:"first_name" 	gorm:"type:varchar(255);not null;"`
	Middlename      string `json:"middle_name" 	gorm:"type:varchar(255);not null;"  `
	Lastname        string `json:"last_name"  	gorm:"type:varchar(255);not null;"`
	Password        string `json:"password"`
	InspectorsCount uint   `json:"inspectors_count" gorm:"default:true"`
	GarageID        uint64 `json:"garage_id"`
}
