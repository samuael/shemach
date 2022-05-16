package pgx_storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/merchant"
)

type MerchantRepo struct {
	DB *pgxpool.Pool
}

func NewMerchantRepo(conn *pgxpool.Pool) merchant.IMerchantRepo {
	return &MerchantRepo{
		DB: conn,
	}
}

// RegisterMerchant ...
func (repo *MerchantRepo) RegisterMerchant(ctx context.Context, Merchant *model.Merchant) (int, error) {
	addressID := 0
	MerchantID := 0
	er := repo.DB.QueryRow(ctx, `select * from createMerchant(cast($1 as varchar),cast($2 as varchar),cast( $3 as varchar),cast($4 as varchar),cast($5 as text) ,cast ($6 as char(3)) ,cast ($7 as int),
	cast ($8 as varchar),cast ($9 as varchar) ,cast ($10 as varchar) ,cast ($11 as varchar) ,cast( $12 as varchar), cast ($13 as varchar),cast ($14 as varchar),cast($15 as varchar))`,
		Merchant.Firstname, Merchant.Lastname, Merchant.Phone, Merchant.Email, Merchant.Password, Merchant.Lang, Merchant.RegisteredBy,
		Merchant.Address.Kebele, Merchant.Address.Woreda, Merchant.Address.City, Merchant.Address.Region, Merchant.Address.Zone,
		Merchant.Address.UniqueAddressName, fmt.Sprint(Merchant.Address.Latitude), fmt.Sprint(Merchant.Address.Longitude),
	).Scan((&MerchantID))
	if er != nil {
		if strings.Contains(er.Error(), "duplicate key value violates unique constraint") {
			return -4, er
		}
		log.Println(er.Error())
		return MerchantID, er
	}
	era := repo.DB.QueryRow(ctx, "select address_id from admin where id=$1", Merchant.ID).Scan(&addressID)
	if era == nil {
		addressID = 0
	}
	if MerchantID < -1 {
		return MerchantID, errors.New("unauthorized")
	} else if MerchantID == -2 {
		return MerchantID, errors.New("unacceptable address information")
	} else if MerchantID == -3 {
		return MerchantID, errors.New("error while creating the Merchant instance")
	}
	Merchant.ID = uint64(MerchantID)
	return int(Merchant.ID), nil
}
func (repo *MerchantRepo) GetMerchantByID(ctx context.Context, id int) (*model.Merchant, error) {
	merchant := &model.Merchant{}
	er := repo.DB.QueryRow(ctx, `select id,firstname,lastname,phone,email,imageurl,created_at,password,lang,stores,posts_count,registerd_by,address_ref  from merchant where id=$1`, id).
		Scan(
			&(merchant.ID), &(merchant.Firstname), &(merchant.Lastname), &(merchant.Phone), &(merchant.Email), &(merchant.Imgurl), &(merchant.CreatedAt), &(merchant.Password), &(merchant.Lang), &(merchant.StoresCount),
			&(merchant.PostsCount), &(merchant.RegisteredBy), &(merchant.AddressRef),
		)
	if er != nil {
		println(er.Error())
		return nil, er
	}
	var address model.Address
	latitude := ""
	longitude := ""
	ers := repo.DB.QueryRow(ctx, `select address_id,kebele,woreda,city,region,unique_name,latitude,zone,longitude from address where address_id=$1`, merchant.AddressRef).
		Scan(&(address.ID), &(address.Kebele), &(address.Woreda), &(address.City), &(address.Region), &(address.UniqueAddressName), &(latitude), &(address.Zone), &(longitude))
	if ers == nil {
		address.Latitude, _ = strconv.ParseFloat(latitude, 64)
		address.Longitude, _ = strconv.ParseFloat(longitude, 64)
		merchant.Address = &address
	}
	return merchant, nil
}
