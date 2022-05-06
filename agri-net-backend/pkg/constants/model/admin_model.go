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
		admin.MerchantsCreated = a.MerchantsCreated.(uint64)
	}
	if a.StorsCreated != nil {
		admin.StoresCreated = a.StorsCreated.(uint64)
	}
	if a.FieldAddressRef != nil {
		admin.FieldAddressRef = a.FieldAddressRef.(int64)
	}
	if a.CreatedBy != nil {
		admin.CreatedBy = a.CreatedBy.(int)
	}
	return admin
}
