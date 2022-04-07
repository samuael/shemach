package model

// Inspector  ...
type Inspector struct {
	ID              uint   `json:"id"  				gorm:"primaryKey;autoIncrement:true" `
	Email           string `json:"email"  			gormSigningKey:"type:varchar(255);not null; unique;"`
	Firstname       string `json:"first_name" 		gorm:"type:varchar(255);not null;"`
	Middlename      string `json:"middle_name" 		gorm:"type:varchar(255);not null;"`
	Lastname        string `json:"last_name" 		gorm:"type:varchar(255);"`
	Password        string `json:"password" 		gorm:"type:varchar(255);not null;"`
	InspectionCount uint   `json:"inspection_count" gorm:"type:integer; default:0;"`
	Imageurl        string `json:"imgurl" 			gorm:"type:varchar(255);"`
	GarageID        uint64 `json:"garage_id"`
	Createdby       uint   `json:"created_by"`
}
