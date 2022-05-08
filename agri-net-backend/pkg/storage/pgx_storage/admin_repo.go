package pgx_storage

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/admin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
)

type AdminRepo struct {
	DB *pgxpool.Pool
}

func NewAdminRepo(conn *pgxpool.Pool) admin.IAdminRepo {
	return &AdminRepo{
		DB: conn,
	}
}

func (repo *AdminRepo) GetAdminByEmail(ctx context.Context, email string) (*model.Admin, error) {
	admin := &model.AdminNullable{}

	er := repo.DB.QueryRow(ctx, `select id,firstname,lastname,phone,email,imageurl,created_at,password,lang,merchants_created,
	stores_created,address_id,created_by from admin where email=$1`, email).
		Scan(&(admin.ID), &(admin.Firstname), &(admin.Lastname), &(admin.Phone), &(admin.Email), &(admin.Imgurl), &(admin.CreatedAt), &(admin.Password), &(admin.Lang), &(admin.MerchantsCreated),
			&(admin.StorsCreated), &(admin.FieldAddressRef), &(admin.CreatedBy),
		)
	if er != nil {
		return nil, er
	}

	latitude := ""
	longitude := ""
	madmin := admin.GetAdmin()
	address := &model.Address{}
	ers := repo.DB.QueryRow(ctx, `select address_id,kebele,woreda,city,region,unique_name,latitude,zone,longitude from address where address_id=$1`, admin.FieldAddressRef).
		Scan(&(address.ID), &(address.Kebele), &(address.Woreda), &(address.City), &(address.Region), &(address.UniqueAddressName), &(latitude), &(address.Zone), &(longitude))
	if ers == nil {
		address.Latitude, _ = strconv.ParseFloat(latitude, 64)
		address.Longitude, _ = strconv.ParseFloat(longitude, 64)
		madmin.FieldAddress = address
	} else {
		println(ers.Error())
	}
	return madmin, er
}
func (repo *AdminRepo) CreateAdmin(ctx context.Context, admin *model.Admin) (int, int, error) {
	addressID := 0
	adminID := 0
	if admin.FieldAddress == nil {
		admin.FieldAddress = &model.Address{}
	}
	er := repo.DB.QueryRow(ctx, `select * from createAdminInstance(cast($1 as varchar),cast($2 as varchar),cast( $3 as varchar),cast($4 as varchar),cast($5 as text) ,cast ($6 as char(3)) ,cast ($7 as int),
	cast ($8 as varchar),cast ($9 as varchar) ,cast ($10 as varchar) ,cast ($11 as varchar) ,cast( $12 as varchar), cast ($13 as varchar),cast ($14 as varchar),cast($15 as varchar))`,
		admin.Firstname, admin.Lastname, admin.Phone, admin.Email, admin.Password, admin.Lang, admin.CreatedBy,
		admin.FieldAddress.Kebele, admin.FieldAddress.Woreda, admin.FieldAddress.City, admin.FieldAddress.Region, admin.FieldAddress.Zone,
		admin.FieldAddress.UniqueAddressName, fmt.Sprint(admin.FieldAddress.Latitude), fmt.Sprint(admin.FieldAddress.Longitude),
	).Scan((&adminID))
	if er != nil {
		return adminID, addressID, er
	}
	era := repo.DB.QueryRow(ctx, "select address_id from admin where id=$1", admin.ID).Scan(&addressID)
	if era == nil {
		addressID = 0
	}
	if adminID < -1 {
		return adminID, 0, errors.New("unauthorized admin instance creation")
	} else if adminID == -2 {
		return adminID, 0, errors.New("error while creating address")
	} else if adminID == -3 {
		return adminID, 0, errors.New("error while creating the admin instance")
	}
	admin.ID = uint64(adminID)
	if admin.FieldAddress != nil {
		admin.FieldAddress.ID = uint(addressID)
	}
	return int(admin.ID), int(admin.FieldAddress.ID), nil
}
func (repo *AdminRepo) GetAdmins(ctx context.Context, offset, limit int) ([]*model.Admin, error) {
	admins := []*model.Admin{}
	rows, ge := repo.DB.Query(ctx, `select id,firstname,lastname,phone,email,imageurl,created_at,password,lang,
	merchants_created,stores_created,address_id,created_by from admin OFFSET $1 LIMIT $2`, offset, limit)
	if ge != nil {
		return admins, ge
	}

	for rows.Next() {
		admin := &model.AdminNullable{}
		er := rows.Scan(&(admin.ID), &(admin.Firstname), &(admin.Lastname), &(admin.Phone), &(admin.Email), &(admin.Imgurl), &(admin.CreatedAt), &(admin.Password), &(admin.Lang), &(admin.MerchantsCreated),
			&(admin.StorsCreated), &(admin.FieldAddressRef), &(admin.CreatedBy),
		)
		if er != nil {
			continue
		}

		latitude := ""
		longitude := ""
		adminM := admin.GetAdmin()
		address := &model.Address{}
		ers := repo.DB.QueryRow(ctx, `select address_id,kebele,woreda,city,region,unique_name,latitude,zone,longitude from address where address_id=$1`, admin.FieldAddressRef).
			Scan(&(address.ID), &(address.Kebele), &(address.Woreda), &(address.City), &(address.Region), &(address.UniqueAddressName), &(latitude), &(address.Zone), &(longitude))
		if ers == nil {
			address.Latitude, _ = strconv.ParseFloat(latitude, 64)
			address.Longitude, _ = strconv.ParseFloat(longitude, 64)
			adminM.FieldAddress = address
		} else {
			println(ers.Error())
		}
		admins = append(admins, adminM)
	}
	return admins, nil
}
func (repo *AdminRepo) DeleteAdminByID(ctx context.Context, id uint64) (int, error) {
	status := 0
	er := repo.DB.QueryRow(ctx, "select * from deleteadminById($1)", id).Scan(&status)
	if er != nil {
		return status, er
	}
	return status, nil
}

func (repo *AdminRepo) GetAdminByID(ctx context.Context, id uint64) (*model.Admin, error) {
	admin := &model.AdminNullable{}

	er := repo.DB.QueryRow(ctx, `select id,firstname,lastname,phone,email,imageurl,created_at,password,lang,merchants_created,
	stores_created,address_id,created_by from admin where id=$1`, id).
		Scan(&(admin.ID), &(admin.Firstname), &(admin.Lastname), &(admin.Phone), &(admin.Email), &(admin.Imgurl), &(admin.CreatedAt), &(admin.Password), &(admin.Lang), &(admin.MerchantsCreated),
			&(admin.StorsCreated), &(admin.FieldAddressRef), &(admin.CreatedBy),
		)
	if er != nil {
		println(er.Error())
		return nil, er
	}
	madmin := admin.GetAdmin()
	address := &model.Address{}
	latitude := ""
	longitude := ""
	ers := repo.DB.QueryRow(ctx, `select address_id,kebele,woreda,city,region,unique_name,latitude,zone,longitude from address where address_id=$1`, admin.FieldAddressRef).
		Scan(&(address.ID), &(address.Kebele), &(address.Woreda), &(address.City), &(address.Region), &(address.UniqueAddressName), &(latitude), &(address.Zone), &(longitude))
	if ers == nil {
		address.Latitude, _ = strconv.ParseFloat(latitude, 64)
		address.Longitude, _ = strconv.ParseFloat(longitude, 64)
		madmin.FieldAddress = address
	} else {
		println(ers.Error())
	}
	return madmin, er
}
