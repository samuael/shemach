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
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
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
	er := repo.DB.QueryRow(ctx, `select id,firstname,lastname,phone,email,imageurl,created_at,password,lang,stores,posts_count,registerd_by,address_ref,subscriptions  from merchant where id=$1`, id).
		Scan(
			&(merchant.ID), &(merchant.Firstname), &(merchant.Lastname), &(merchant.Phone), &(merchant.Email), &(merchant.Imgurl), &(merchant.CreatedAt), &(merchant.Password), &(merchant.Lang), &(merchant.StoresCount),
			&(merchant.PostsCount), &(merchant.RegisteredBy), &(merchant.AddressRef), &(merchant.Subscriptions),
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

func (repo *MerchantRepo) CreateSubscriptions(ctx context.Context, productid uint8, merchantid uint64) (status int) {
	status = 0
	if er := repo.DB.QueryRow(ctx, "select * from MerchantCreateNewSubscription( $1,$2::smallint )", merchantid, productid).Scan(&(status)); er != nil {
		return state.STATUS_DBQUERY_ERROR
	}
	return status
}

func (repo *MerchantRepo) UnsubscribeProduct(ctx context.Context, productid uint8, merchantid uint64) (status int) {
	status = 0
	if er := repo.DB.QueryRow(ctx, "select * from MerchantUnSubscribeToProduct( $1,$2::smallint )", merchantid, productid).Scan(&(status)); er != nil {
		return state.STATUS_DBQUERY_ERROR
	}
	return status
}

func (repo *MerchantRepo) SearchMerchants(ctx context.Context, phone, name string, createdBy uint64, offset, limit uint) ([]*model.Merchant, error) {
	merchants := []*model.Merchant{}
	values := []interface{}{}
	count := 1
	statement := `select id,firstname,lastname,phone,email,imageurl,created_at,password,lang,stores,posts_count,registerd_by,address_ref,subscriptions from merchant where `
	if phone != "" {
		statement = fmt.Sprintf("%s phone ILIKE $%d", statement, count)
		values = append(values, "%"+strings.Trim(phone, " +")+"%")
		count++
	}
	name = strings.Trim(name, " ")
	if len(strings.Split(name, " ")) > 1 {
		if count > 1 {
			statement = fmt.Sprintf(" %s or ", statement)
		}
		statement = fmt.Sprintf(" %s ( firstname ILIKE $"+strconv.Itoa(count), statement)
		values = append(values, "%"+(strings.Split(name, " ")[0])+"%")
		statement = fmt.Sprintf("%s and ", statement)
		count++
		statement = fmt.Sprintf("%s lastname ILIKE $%d ) ", statement, count)
		values = append(values, "%"+(strings.Split(name, " ")[1])+"%")
		count++
	} else if name != "" {
		if count > 1 {
			statement = fmt.Sprintf(" %s or ", statement)
		}
		statement = fmt.Sprintf("%s firstname ILIKE $"+strconv.Itoa(count), statement)
		values = append(values, "%"+name+"%")
		statement = fmt.Sprintf("%s or ", statement)
		count++
		statement = fmt.Sprintf("%s lastname ILIKE $"+strconv.Itoa(count), statement)
		values = append(values, "%"+name+"%")
		count++
	}
	if createdBy > 0 {
		if count > 1 {
			statement = fmt.Sprintf(" %s or ", statement)
		}
		statement = fmt.Sprintf(" %s registerd_by=$"+strconv.Itoa(int(count)), statement)
		values = append(values, createdBy)
		count++
	}
	statement = fmt.Sprintf("%s ORDER BY id DESC OFFSET $%d LIMIT $%d ", statement, count, count+1)
	values = append(values, offset, limit)
	rows, er := repo.DB.Query(ctx, statement, values...)
	if er != nil {

		return nil, er
	}
	for rows.Next() {
		merchant := &model.Merchant{}
		erf := rows.Scan(
			&(merchant.ID), &(merchant.Firstname), &(merchant.Lastname), &(merchant.Phone), &(merchant.Email), &(merchant.Imgurl), &(merchant.CreatedAt), &(merchant.Password), &(merchant.Lang), &(merchant.StoresCount),
			&(merchant.PostsCount), &(merchant.RegisteredBy), &(merchant.AddressRef), &(merchant.Subscriptions),
		)
		if erf != nil {
			println(erf.Error())
			continue
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
		merchants = append(merchants, merchant)
	}
	return merchants, nil
}
