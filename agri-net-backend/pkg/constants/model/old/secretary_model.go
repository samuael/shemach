package model

// Secretary ...
type Secretary struct {
	ID         uint   `json:"id"          gorm:"primaryKey;autoIncrement:true" `
	Email      string `json:"email"       gormSigningKey:"type:varchar(255);not null; unique;"`
	Firstname  string `json:"first_name"  gorm:"type:varchar(255);not null;"`
	Middlename string `json:"middle_name" gorm:"type:varchar(255);not null;"`
	Lastname   string `json:"last_name"   gorm:"type:varchar(255);not null;"`
	Password   string `json:"password"`
	Createdby  uint   `json:"created_by"`
	GarageID   uint64 `json:"garage_id"`
}
