package model

type Admin struct {
	User
	MerchantsCreated uint64   `json:"merchants_created"`
	StoresCreated    uint64   `json:"stores_created"`
	FieldAddress     *Address `json:"address"`
	CreatedBy        int      `json:"created_by"`
	FieldAddressRef  int64    `json:"-"`
}

type AdminNullable struct {
	User
	MerchantsCreated interface{} `json:"merchants_created"`
	StorsCreated     interface{} `json:"stores_created"`
	FieldAddressRef  interface{} `json:"address"`
	CreatedBy        interface{} `json:"created_by"`
}

func (a *AdminNullable) GetAdmin() *Admin {
	admin := &Admin{
		User: a.User,
	}
	if a.MerchantsCreated != nil {
		admin.MerchantsCreated = uint64(a.MerchantsCreated.(int32))
	}
	if a.StorsCreated != nil {
		admin.StoresCreated = uint64(a.StorsCreated.(int32))
	}
	if a.FieldAddressRef != nil {
		admin.FieldAddressRef = int64(a.FieldAddressRef.(int32))
	}
	if a.CreatedBy != nil {
		admin.CreatedBy = int(a.CreatedBy.(int32))
	}
	return admin
}
